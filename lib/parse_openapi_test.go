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
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOpenAPI(t *testing.T) {
	t.Run("ParseOpenAPI", func(t *testing.T) {
		o := NewOpenAPI("../testdata/openapi.json")
		res, err := o.Parse(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, "/pets", res[0].Path())

		assert.Equal(t, "GET", res[0].method[0].method)
		assert.Equal(t, "POST", res[0].method[1].method)
		if res[0].method[1].bodys != nil {
			by, _ := res[0].method[1].bodys[0].MarshalJSON()
			fmt.Println(string(by))
		}

		assert.Equal(t, "/pets/{id}", res[1].Path())

		assert.Equal(t, "DELETE", res[1].method[0].method)
		assert.Equal(t, "GET", res[1].method[1].method)
	})
}
