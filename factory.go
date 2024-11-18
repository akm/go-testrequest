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

func (f *Factory) New(method string, options ...Option) Func {
	return func(t *testing.T) *http.Request {
		t.Helper()
		b := &builder{method: method, headers: make(http.Header)}
		for _, option := range append(f.defaultOptions, options...) {
			option(b)
		}
		return b.build(t)
	}
}

func (f *Factory) GET(v ...Option) Func {
	return f.New(http.MethodGet, v...)
}

func (f *Factory) HEAD(v ...Option) Func {
	return f.New(http.MethodHead, v...)
}

func (f *Factory) POST(v ...Option) Func {
	return f.New(http.MethodPost, v...)
}

func (f *Factory) PUT(v ...Option) Func {
	return f.New(http.MethodPut, v...)
}

func (f *Factory) PATCH(v ...Option) Func {
	return f.New(http.MethodPatch, v...)
}

func (f *Factory) DELETE(v ...Option) Func {
	return f.New(http.MethodDelete, v...)
}

func (f *Factory) CONNECT(v ...Option) Func {
	return f.New(http.MethodConnect, v...)
}

func (f *Factory) OPTIONS(v ...Option) Func {
	return f.New(http.MethodOptions, v...)
}

func (f *Factory) TRACE(v ...Option) Func {
	return f.New(http.MethodTrace, v...)
}

var DefaultFactory = NewFactory()
