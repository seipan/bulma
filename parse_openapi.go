package bluma

import (
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

func (o *OpenAPI) Parse() {
	doc, err := openapi3.NewLoader().LoadFromFile(o.Pass())
	if err != nil {
		fmt.Println(err)
	}
}
