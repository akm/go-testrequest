package testrequest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Option = func(*builder)

func Context(v context.Context) Option { return func(b *builder) { b.context = v } }
func BaseUrl(v string) Option          { return func(b *builder) { b.baseUrl = v } }
func Scheme(v string) Option           { return func(b *builder) { b.scheme = v } }
func Host(v string) Option             { return func(b *builder) { b.host = v } }
func Port(v int) Option                { return func(b *builder) { b.port = v } }

func Path(v string, args ...interface{}) Option {
	return func(b *builder) { b.path = fmt.Sprintf(v, args...) }
}
func Query(k, v string) Option {
	return func(b *builder) { b.query = append(b.query, [2]string{k, v}) }
}
func Header(k, v string) Option { return func(b *builder) { b.headers.Add(k, v) } }
func Cookie(v *http.Cookie) Option {
	return func(b *builder) { b.cookies = append(b.cookies, v) }
}

func Body(v *io.Reader) Option { return func(b *builder) { b.body = *v } }
func BodyString(v string) Option {
	return func(b *builder) { b.body = io.NopCloser(strings.NewReader(v)) }
}
func BodyBytes(v []byte) Option {
	return func(b *builder) { b.body = io.NopCloser(bytes.NewReader(v)) }
}
