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

package lib

import (
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

type Attacker struct {
	Path        Path
	MethodIndex int
	Body        []byte
	Header      http.Header
	Frequency   int
	Duration    time.Duration
}

func (atk *Attacker) Attack() vegeta.Metrics {
	target := vegeta.Target{
		Method: atk.Path.method[atk.MethodIndex].method,
		URL:    atk.Path.path,
		Body:   atk.Body,
		Header: atk.Header,
	}
	targeter := vegeta.NewStaticTargeter(target)
	rate := vegeta.Rate{Freq: atk.Frequency, Per: time.Second}

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, atk.Duration, "Vegeta Load Testing") {
		metrics.Add(res)
	}
	metrics.Close()

	return metrics
}
