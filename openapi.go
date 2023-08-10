package bluma

import "github.com/getkin/kin-openapi/openapi3"

type Method struct {
	method string
	params openapi3.Parameters
	bodys  []*openapi3.SchemaRef
}

type Path struct {
	path   string
	method []Method
}
