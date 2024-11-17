package testrequest

import "net/http"

func New(method string, options ...Option) *http.Request {
	b := &builder{method: method, headers: make(http.Header)}
	for _, option := range options {
		option(b)
	}
	return b.build()
}

func GET(v ...Option) *http.Request     { return New(http.MethodGet, v...) }
func HEAD(v ...Option) *http.Request    { return New(http.MethodHead, v...) }
func POST(v ...Option) *http.Request    { return New(http.MethodPost, v...) }
func PUT(v ...Option) *http.Request     { return New(http.MethodPut, v...) }
func PATCH(v ...Option) *http.Request   { return New(http.MethodPatch, v...) }
func DELETE(v ...Option) *http.Request  { return New(http.MethodDelete, v...) }
func CONNECT(v ...Option) *http.Request { return New(http.MethodConnect, v...) }
func OPTIONS(v ...Option) *http.Request { return New(http.MethodOptions, v...) }
func TRACE(v ...Option) *http.Request   { return New(http.MethodTrace, v...) }
