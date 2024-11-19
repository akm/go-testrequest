package testrequest

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type builder struct {
	context context.Context
	method  string
	baseUrl string
	scheme  string
	host    string
	port    string
	path    string
	query   [][2]string
	headers http.Header
	cookies []*http.Cookie
	body    io.Reader
}

func (b *builder) build() (*http.Request, error) {
	url := b.buildURL()
	req, err := http.NewRequest(b.method, url, b.body)
	if err != nil {
		return nil, err
	}
	if b.context != nil {
		req = req.WithContext(b.context)
	}
	req.Header = b.headers
	for _, cookie := range b.cookies {
		req.AddCookie(cookie)
	}
	return req, nil
}

func (b *builder) buildURL() string {
	var url string
	if b.baseUrl != "" {
		url = b.baseUrl
	} else {
		scheme := "http"
		host := "localhost"
		port := "80"
		if b.scheme != "" {
			scheme = b.scheme
		}
		if b.host != "" {
			host = b.host
		}
		if b.port != "" {
			port = b.port
		}
		url = fmt.Sprintf("%s://%s:%s", scheme, host, port)
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
