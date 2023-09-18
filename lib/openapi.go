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

import "github.com/getkin/kin-openapi/openapi3"

type Body struct {
	name  string
	shema *openapi3.SchemaRef
}

type Method struct {
	method string
	params openapi3.Parameters
	bodys  []Body
}

func (m *Method) Method() string {
	return m.method
}

func (m *Method) Params() openapi3.Parameters {
	return m.params
}

func (m *Method) Bodys() []Body {
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
