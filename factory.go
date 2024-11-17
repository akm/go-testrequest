package testrequest

import (
	"net/http"
	"testing"
)

type Factory struct {
	defaultOptions []Option
}

func NewFactory(defaultOptions ...Option) *Factory {
	return &Factory{defaultOptions: defaultOptions}
}

func (f *Factory) New(t *testing.T, method string, options ...Option) *http.Request {
	t.Helper()
	b := &builder{method: method, headers: make(http.Header)}
	for _, option := range append(f.defaultOptions, options...) {
		option(b)
	}
	return b.build(t)
}

func (f *Factory) GET(t *testing.T, v ...Option) *http.Request {
	return f.New(t, http.MethodGet, v...)
}

func (f *Factory) HEAD(t *testing.T, v ...Option) *http.Request {
	return f.New(t, http.MethodHead, v...)
}

func (f *Factory) POST(t *testing.T, v ...Option) *http.Request {
	return f.New(t, http.MethodPost, v...)
}

func (f *Factory) PUT(t *testing.T, v ...Option) *http.Request {
	return f.New(t, http.MethodPut, v...)
}

func (f *Factory) PATCH(t *testing.T, v ...Option) *http.Request {
	return New(t, http.MethodPatch, v...)
}

func (f *Factory) DELETE(t *testing.T, v ...Option) *http.Request {
	return New(t, http.MethodDelete, v...)
}

func (f *Factory) CONNECT(t *testing.T, v ...Option) *http.Request {
	return New(t, http.MethodConnect, v...)
}

func (f *Factory) OPTIONS(t *testing.T, v ...Option) *http.Request {
	return New(t, http.MethodOptions, v...)
}

func (f *Factory) TRACE(t *testing.T, v ...Option) *http.Request {
	return New(t, http.MethodTrace, v...)
}

var DefaultFactory = NewFactory()
