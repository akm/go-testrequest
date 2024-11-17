package testrequest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type builder struct {
	method  string
	baseUrl string
	scheme  string
	host    string
	port    int
	path    string
	query   [][2]string
	headers http.Header
	cookies []*http.Cookie
	body    io.Reader
}

func (b *builder) build() *http.Request {
	url := b.buildURL()
	req := httptest.NewRequest(b.method, url, b.body)
	req.Header = b.headers
	for _, cookie := range b.cookies {
		req.AddCookie(cookie)
	}
	return req
}

func (b *builder) buildURL() string {
	var url string
	if b.baseUrl != "" {
		url = b.baseUrl
	} else if b.scheme != "" || b.host != "" || b.port != 0 {
		scheme := "http"
		host := "localhost"
		port := 80
		if b.scheme != "" {
			scheme = b.scheme
		}
		if b.host != "" {
			host = b.host
		}
		if b.port != 0 {
			port = b.port
		}
		url = fmt.Sprintf("%s://%s:%d", scheme, host, port)
	}

	url += b.path
	if len(b.query) > 0 {
		url += "?"
		for i, q := range b.query {
			if i > 0 {
				url += "&"
			}
			url += q[0] + "=" + q[1]
		}
	}
	return url
}
