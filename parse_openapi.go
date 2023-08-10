package bluma

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

func (o *OpenAPI) Parse(ctx context.Context) {
	doc, err := openapi3.NewLoader().LoadFromFile(o.Pass())
	if err != nil {
		fmt.Println(err)
	}

	_ = doc.Validate(ctx)
	for _, pathItem := range doc.Paths.InMatchingOrder() {
		oprs := doc.Paths.Find(pathItem).Operations()

		var paths []Method

		for mtd, opr := range oprs {
			var params openapi3.Parameters
			var bodys []*openapi3.SchemaRef

			if opr.Parameters != nil {
				for _, param := range opr.Parameters {
					params = append(params, param)
				}
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
			paths = append(paths, path)
		}
	}
}
