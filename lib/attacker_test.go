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
	"fmt"
	"testing"
)

func TestAttack(t *testing.T) {
	t.Run("Attack", func(t *testing.T) {
		atk := Attacker{
			Path: Path{
				path: "http://localhost:8080/health",
				method: []Method{
					{
						method: "GET",
					},
				},
			},
			MethodIndex: 0,
			Frequency:   110,
			Duration:    2,
		}
		metrics := atk.Attack()
		fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
		fmt.Printf("max percentile: %s\n", metrics.Latencies.Max)
		fmt.Printf("mean percentile: %s\n", metrics.Latencies.Mean)
		fmt.Printf("total percentile: %s\n", metrics.Latencies.Total)

		fmt.Printf("status code: %v\n", metrics.StatusCodes)
	})
}
