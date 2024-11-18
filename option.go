package testrequest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Option is a function that modifies the builder.
type Option = func(*builder)

// Context sets the context for the request.
func Context(v context.Context) Option { return func(b *builder) { b.context = v } }

// BaseUrl sets the base URL for the request.
func BaseUrl(v string) Option { return func(b *builder) { b.baseUrl = v } }

// Scheme sets the URL scheme (http or https) for the request.
func Scheme(v string) Option { return func(b *builder) { b.scheme = v } }

// Host sets the host for the request.
func Host(v string) Option { return func(b *builder) { b.host = v } }

// PortString sets the port for the request as a string.
func PortString(v string) Option { return func(b *builder) { b.port = v } }

// Port sets the port for the request as an integer.
func Port(v int) Option { return func(b *builder) { b.port = strconv.Itoa(v) } }

// Path sets the path for the request, with optional formatting arguments.
func Path(v string, args ...interface{}) Option {
	return func(b *builder) { b.path = fmt.Sprintf(v, args...) }
}

// Query adds a query parameter to the request.
func Query(k, v string) Option {
	return func(b *builder) { b.query = append(b.query, [2]string{k, v}) }
}

// Header adds a header to the request.
func Header(k, v string) Option { return func(b *builder) { b.headers.Add(k, v) } }

// Cookie adds a cookie to the request.
func Cookie(v *http.Cookie) Option {
	return func(b *builder) { b.cookies = append(b.cookies, v) }
}

// Body sets the body of the request from an io.Reader.
func Body(v *io.Reader) Option { return func(b *builder) { b.body = *v } }

// BodyString sets the body of the request from a string.
func BodyString(v string) Option {
	return func(b *builder) { b.body = io.NopCloser(strings.NewReader(v)) }
}

// BodyBytes sets the body of the request from a byte slice.
func BodyBytes(v []byte) Option {
	return func(b *builder) { b.body = io.NopCloser(bytes.NewReader(v)) }
}
