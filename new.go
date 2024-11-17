package testrequest

import (
	"net/http"
	"testing"
)

func New(t *testing.T, method string, options ...Option) *http.Request {
	return DefaultFactory.New(t, method, options...)
}

func GET(t *testing.T, v ...Option) *http.Request     { return DefaultFactory.GET(t, v...) }
func HEAD(t *testing.T, v ...Option) *http.Request    { return DefaultFactory.HEAD(t, v...) }
func POST(t *testing.T, v ...Option) *http.Request    { return DefaultFactory.POST(t, v...) }
func PUT(t *testing.T, v ...Option) *http.Request     { return DefaultFactory.PUT(t, v...) }
func PATCH(t *testing.T, v ...Option) *http.Request   { return DefaultFactory.PATCH(t, v...) }
func DELETE(t *testing.T, v ...Option) *http.Request  { return DefaultFactory.DELETE(t, v...) }
func CONNECT(t *testing.T, v ...Option) *http.Request { return DefaultFactory.CONNECT(t, v...) }
func OPTIONS(t *testing.T, v ...Option) *http.Request { return DefaultFactory.OPTIONS(t, v...) }
func TRACE(t *testing.T, v ...Option) *http.Request   { return DefaultFactory.TRACE(t, v...) }
