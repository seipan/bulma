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

	"github.com/getkin/kin-openapi/openapi3"
)

type OpenAPI struct {
	pass string
}

func NewOpenAPI(pass string) *OpenAPI {
	return &OpenAPI{
		pass: pass,
	}
}

func (o *OpenAPI) Pass() string {
	return o.pass
}

func (o *OpenAPI) Parse(ctx context.Context) ([]Path, error) {
	doc, err := openapi3.NewLoader().LoadFromFile(o.Pass())
	if err != nil {
		return nil, fmt.Errorf("failed to load openapi.yaml: %w", err)
	}

	var paths []Path

	_ = doc.Validate(ctx)
	for _, pathItem := range doc.Paths.InMatchingOrder() {
		oprs := doc.Paths.Find(pathItem).Operations()

		var mtds []Method

		for mtd, opr := range oprs {
			var params openapi3.Parameters
			var bodys []*openapi3.SchemaRef

			if opr.Parameters != nil {
				params = append(params, opr.Parameters...)
			}

			if opr.RequestBody != nil {
				for _, param := range opr.RequestBody.Value.Content["application/json"].Schema.Value.Properties {
					bodys = append(bodys, param)
				}
			}

			path := Method{
				method: mtd,
				params: params,
				bodys:  bodys,
			}
			mtds = append(mtds, path)
		}

		path := Path{
			path:   pathItem,
			method: mtds,
		}
		paths = append(paths, path)

	}
	return paths, nil
}
