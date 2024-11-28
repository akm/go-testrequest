package testrequest

import (
	"net/http"
	"testing"
)

// New creates a new request with the specified method and options.
func New(method string, options ...Option) Func {
	b := NewBuilder(method)
	// Apply default options and additional options to the builder.
	for _, option := range options {
		option(b)
	}
	return func(t *testing.T) *http.Request {
		t.Helper()
		// Build and return the HTTP request.
		req, err := b.build()
		if err != nil {
			t.Fatal(err)
		}
		return req
	}
}

// GET creates a new GET request with the specified options.
func GET(v ...Option) Func { return New(http.MethodGet, v...) }

// HEAD creates a new HEAD request with the specified options.
func HEAD(v ...Option) Func { return New(http.MethodHead, v...) }

// POST creates a new POST request with the specified options.
func POST(v ...Option) Func { return New(http.MethodPost, v...) }

// PUT creates a new PUT request with the specified options.
func PUT(v ...Option) Func { return New(http.MethodPut, v...) }

// PATCH creates a new PATCH request with the specified options.
func PATCH(v ...Option) Func { return New(http.MethodPatch, v...) }

// DELETE creates a new DELETE request with the specified options.
func DELETE(v ...Option) Func { return New(http.MethodDelete, v...) }

// CONNECT creates a new CONNECT request with the specified options.
func CONNECT(v ...Option) Func { return New(http.MethodConnect, v...) }

// OPTIONS creates a new OPTIONS request with the specified options.
func OPTIONS(v ...Option) Func { return New(http.MethodOptions, v...) }

// TRACE creates a new TRACE request with the specified options.
func TRACE(v ...Option) Func { return New(http.MethodTrace, v...) }
