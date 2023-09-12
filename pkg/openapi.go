package pkg

import "github.com/getkin/kin-openapi/openapi3"

type Method struct {
	method string
	params openapi3.Parameters
	bodys  []*openapi3.SchemaRef
}

func (m *Method) Method() string {
	return m.method
}

func (m *Method) Params() openapi3.Parameters {
	return m.params
}

func (m *Method) Bodys() []*openapi3.SchemaRef {
	return m.bodys
}

type Path struct {
	path   string
	method []Method
}

func (p *Path) Path() string {
	return p.path
}

func (p *Path) Method(index int) Method {
	return p.method[index]
}
