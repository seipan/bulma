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

		for mtd, opr := range oprs {
			fmt.Println(mtd, opr)
		}
	}
}
