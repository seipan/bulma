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
	"time"

	"github.com/seipan/bulma/lib"
)

func ParseAndAttack(ctx context.Context, path string, freq int, duration time.Duration) error {
	oapi := lib.NewOpenAPI(path)
	paths, err := oapi.Parse(ctx)
	if err != nil {
		return fmt.Errorf("failed to parse openapi: %w", err)
	}
	atks, err := ParthOpenAPItoAttacker(paths, freq, duration)
	if err != nil {
		return fmt.Errorf("failed to convert openapi to attacker: %w", err)
	}
	for _, atk := range atks {
		atk.Attack()
	}
	return nil
}

func ParthOpenAPItoAttacker(pathes []lib.Path, freq int, duration time.Duration) ([]lib.Attacker, error) {
	var res []lib.Attacker
	for i, path := range pathes {
		mtd := path.Method(0)
		bodys := mtd.Bodys()
		body, err := createBody(bodys)
		if err != nil {
			return nil, err
		}
		atk := lib.Attacker{
			Path:        path,
			MethodIndex: i,
			Body:        body,
			Header:      http.Header{},
			Frequency:   freq,
			Duration:    duration,
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
