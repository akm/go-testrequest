package testrequest

import (
	"net/http"
	"testing"
)

// Factory is a struct that holds default options for creating HTTP requests.
type Factory struct {
	defaultOptions []Option
}

// NewFactory creates a new Factory with the given default options.
func NewFactory(defaultOptions ...Option) *Factory {
	return &Factory{defaultOptions: defaultOptions}
}

// New creates a new HTTP request function with the specified method and options.
func (f *Factory) New(method string, options ...Option) Func {
	return func(t *testing.T) *http.Request {
		t.Helper()
		b := &builder{method: method, headers: make(http.Header)}
		// Apply default options and additional options to the builder.
		for _, option := range append(f.defaultOptions, options...) {
			option(b)
		}
		// Build and return the HTTP request.
		return b.build(t)
	}
}

// GET creates a new HTTP GET request function with the specified options.
func (f *Factory) GET(v ...Option) Func {
	return f.New(http.MethodGet, v...)
}

// HEAD creates a new HTTP HEAD request function with the specified options.
func (f *Factory) HEAD(v ...Option) Func {
	return f.New(http.MethodHead, v...)
}

// POST creates a new HTTP POST request function with the specified options.
func (f *Factory) POST(v ...Option) Func {
	return f.New(http.MethodPost, v...)
}

// PUT creates a new HTTP PUT request function with the specified options.
func (f *Factory) PUT(v ...Option) Func {
	return f.New(http.MethodPut, v...)
}

// PATCH creates a new HTTP PATCH request function with the specified options.
func (f *Factory) PATCH(v ...Option) Func {
	return f.New(http.MethodPatch, v...)
}

// DELETE creates a new HTTP DELETE request function with the specified options.
func (f *Factory) DELETE(v ...Option) Func {
	return f.New(http.MethodDelete, v...)
}

// CONNECT creates a new HTTP CONNECT request function with the specified options.
func (f *Factory) CONNECT(v ...Option) Func {
	return f.New(http.MethodConnect, v...)
}

// OPTIONS creates a new HTTP OPTIONS request function with the specified options.
func (f *Factory) OPTIONS(v ...Option) Func {
	return f.New(http.MethodOptions, v...)
}

// TRACE creates a new HTTP TRACE request function with the specified options.
func (f *Factory) TRACE(v ...Option) Func {
	return f.New(http.MethodTrace, v...)
}

var DefaultFactory = NewFactory()
