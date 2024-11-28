package testrequest

import (
	"context"
	"io"
	"net/http"

	"github.com/akm/go-testrequest/builder"
)

// Option is a function that modifies the builder.
type Option = builder.Option

// Context sets the context for the request.
func Context(v context.Context) Option { return builder.Context(v) }

// BaseUrl sets the base URL for the request.
func BaseUrl(v string) Option { return builder.BaseUrl(v) }

// Scheme sets the URL scheme (http or https) for the request.
func Scheme(v string) Option { return builder.Scheme(v) }

// Host sets the host for the request.
func Host(v string) Option { return builder.Host(v) }

// PortString sets the port for the request as a string.
func PortString(v string) Option { return builder.PortString(v) }

// Port sets the port for the request as an integer.
func Port(v int) Option { return builder.Port(v) }

// Path sets the path for the request, with optional formatting arguments.
func Path(v string, args ...interface{}) Option { return builder.Path(v, args...) }

// Query adds a query parameter to the request.
func Query(k, v string) Option { return builder.Query(k, v) }

// Header adds a header to the request.
func Header(k, v string) Option { return builder.Header(k, v) }

// Cookie adds a cookie to the request.
func Cookie(v *http.Cookie) Option { return builder.Cookie(v) }

// Body sets the body of the request from an io.Reader.
func Body(v *io.Reader) Option { return builder.Body(v) }

// BodyString sets the body of the request from a string.
func BodyString(v string) Option { return builder.BodyString(v) }

// BodyBytes sets the body of the request from a byte slice.
func BodyBytes(v []byte) Option { return builder.BodyBytes(v) }
