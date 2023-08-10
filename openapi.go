package bluma

import "github.com/getkin/kin-openapi/openapi3"

type Path struct {
	method string
	params openapi3.Parameters
	bodys  []*openapi3.SchemaRef
}
