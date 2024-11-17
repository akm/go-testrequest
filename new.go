package testrequest

import (
	"net/http"
	"testing"
)

func New(t *testing.T, method string, options ...Option) *http.Request {
	t.Helper()
	b := &builder{method: method, headers: make(http.Header)}
	for _, option := range options {
		option(b)
	}
	return b.build(t)
}

func GET(t *testing.T, v ...Option) *http.Request     { return New(t, http.MethodGet, v...) }
func HEAD(t *testing.T, v ...Option) *http.Request    { return New(t, http.MethodHead, v...) }
func POST(t *testing.T, v ...Option) *http.Request    { return New(t, http.MethodPost, v...) }
func PUT(t *testing.T, v ...Option) *http.Request     { return New(t, http.MethodPut, v...) }
func PATCH(t *testing.T, v ...Option) *http.Request   { return New(t, http.MethodPatch, v...) }
func DELETE(t *testing.T, v ...Option) *http.Request  { return New(t, http.MethodDelete, v...) }
func CONNECT(t *testing.T, v ...Option) *http.Request { return New(t, http.MethodConnect, v...) }
func OPTIONS(t *testing.T, v ...Option) *http.Request { return New(t, http.MethodOptions, v...) }
func TRACE(t *testing.T, v ...Option) *http.Request   { return New(t, http.MethodTrace, v...) }
