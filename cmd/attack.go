// MIT License

// Copyright (c) 2023 Yamasaki Shotaro

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/seipan/bulma/lib"
	vegeta "github.com/tsenart/vegeta/lib"
)

func ParseAndAttack(ctx context.Context, beseEndpoint string, path string, freq int, duration time.Duration) error {
	oapi := lib.NewOpenAPI(path)
	paths, err := oapi.Parse(ctx)
	if err != nil {
		return fmt.Errorf("failed to parse openapi: %w", err)
	}
	atks, err := ParthOpenAPItoAttacker(paths, beseEndpoint, freq, duration)
	if err != nil {
		return fmt.Errorf("failed to convert openapi to attacker: %w", err)
	}
	fmt.Println("--------------------------bulma attack start-------------------------------")
	var wg sync.WaitGroup
	for _, atk := range atks {
		wg.Add(1)
		go func() {
			metrics := atk.Attack()
			outputMetrics(metrics, &atk)
		}()
		wg.Done()
	}
	fmt.Println("--------------------------bulma attack finish-------------------------------")
	wg.Wait()
	return nil
}

func ParthOpenAPItoAttacker(pathes []lib.Path, beseEndpoint string, freq int, duration time.Duration) ([]lib.Attacker, error) {
	var res []lib.Attacker
	for i, path := range pathes {
		mtd := path.Method(0)
		bodys := mtd.Bodys()
		body, err := createBody(bodys)
		if err != nil {
			return nil, err
		}
		path.SetPath(beseEndpoint + path.Path())
		atk := lib.Attacker{
			Path:        path,
			MethodIndex: i,
			Body:        body,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Frequency: freq,
			Duration:  duration,
		}
		res = append(res, atk)
	}
	return res, nil
}

func createBody(bodys []lib.Body) ([]byte, error) {
	mp := make(map[string]interface{})
	for _, body := range bodys {
		mp[body.Name] = body.Shema.Value.Example
	}
	jsonData, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func outputMetrics(metrics vegeta.Metrics, atk *lib.Attacker) {
	fmt.Printf("--------------------------vegeta attack to %s--------------------------\n", atk.Path.Path())
	mtd := atk.Path.Method(atk.MethodIndex)
	fmt.Printf("vegeta attack to method: %s\n", mtd.Method())
	fmt.Printf("path StatusCode: %v\n", metrics.StatusCodes)

	fmt.Println()

	fmt.Printf("max percentile: %s\n", metrics.Latencies.Max)
	fmt.Printf("mean percentile: %s\n", metrics.Latencies.Mean)
	fmt.Printf("total percentile: %s\n", metrics.Latencies.Total)
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)

	fmt.Println()

	fmt.Printf(" earliest: %v\n", metrics.Earliest)
	fmt.Printf(" latest: %v\n", metrics.Latest)

	fmt.Println("-----------------------------------------------------------------------")
}
