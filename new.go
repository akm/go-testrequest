package testrequest

import "net/http"

func New(method string, options ...Option) *http.Request {
	b := &builder{method: method, headers: make(http.Header)}
	for _, option := range options {
		option(b)
	}
	return b.build()
}
