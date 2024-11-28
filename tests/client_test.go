package testrequest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/akm/reqb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientWithServer(t *testing.T) {
	testServer := startEchoServer(t)
	testServer.Start()
	defer testServer.Close()

	testServerURL, err := url.Parse(testServer.URL)
	require.NoError(t, err)

	baseURL := testServer.URL
	baseOpts := reqb.Options{reqb.BaseUrl(baseURL)}

	defaultHeader := func() http.Header {
		return http.Header{
			"Accept-Encoding": []string{"gzip"},
			"User-Agent":      []string{"Go-http-client/1.1"},
		}
	}
	mergeHeader := func(h1, h2 http.Header) http.Header {
		for k, v := range h2 {
			h1[k] = v
		}
		return h1
	}
	expectedHeader := func(h http.Header) http.Header {
		return mergeHeader(defaultHeader(), h)
	}

	type pattern *struct {
		name     string
		funcs    map[string]reqb.Func
		expected *request
	}
	patterns := []pattern{
		{
			"GET /",
			map[string]reqb.Func{
				"ad-hoc":           reqb.GET(reqb.BaseUrl(baseURL)),
				"package function": reqb.GET(baseOpts...),
				"options method":   baseOpts.GET(),
			},
			&request{
				Method: http.MethodGet,
				Url:    "/",
				Header: expectedHeader(http.Header{}),
				Body:   "",
			},
		},
		{
			"POST /users",
			map[string]reqb.Func{
				"ad-hoc": reqb.POST(
					reqb.BaseUrl(baseURL),
					reqb.Path("/users"),
					reqb.BodyString("hello, world"),
				),
				"package function": reqb.POST(
					append(baseOpts,
						reqb.Path("/users"),
						reqb.BodyString("hello, world"),
					)...,
				),
				"options method": baseOpts.POST(
					reqb.Path("/users"),
					reqb.BodyString("hello, world"),
				),
			},
			&request{
				Method: http.MethodPost,
				Url:    "/users",
				Header: expectedHeader(http.Header{}),
				Body:   "hello, world",
			},
		},
		{
			"PUT /users/123",
			map[string]reqb.Func{
				"ad-hoc": reqb.PUT(
					reqb.BaseUrl(baseURL),
					reqb.Path("/users/%d", 123),
					reqb.BodyString("{\"name\":\"foo\"}"),
					reqb.Header("Content-Type", "application/json"),
				),
				"package function": reqb.PUT(
					append(baseOpts,
						reqb.Path("/users/%d", 123),
						reqb.BodyString("{\"name\":\"foo\"}"),
						reqb.Header("Content-Type", "application/json"),
					)...,
				),
				"options method": baseOpts.PUT(
					reqb.Path("/users/%d", 123),
					reqb.BodyString("{\"name\":\"foo\"}"),
					reqb.Header("Content-Type", "application/json"),
				),
			},
			&request{
				Method: http.MethodPut,
				Url:    "/users/123",
				Header: expectedHeader(http.Header{
					"Content-Type": []string{"application/json"},
				}),
				Body: "{\"name\":\"foo\"}",
			},
		},
		{
			"PATCH /users/123",
			map[string]reqb.Func{
				"ad-hoc": reqb.PATCH(
					reqb.BaseUrl(baseURL),
					reqb.Path("/users/%d", 123),
					reqb.BodyBytes([]byte("{\"name\":\"bar\"}")),
					reqb.Header("Content-Type", "application/json"),
					reqb.Cookie(&http.Cookie{Name: "session", Value: "session1"}),
				),
				"package function": reqb.PATCH(
					append(baseOpts,
						reqb.Path("/users/%d", 123),
						reqb.BodyBytes([]byte("{\"name\":\"bar\"}")),
						reqb.Header("Content-Type", "application/json"),
						reqb.Cookie(&http.Cookie{Name: "session", Value: "session1"}),
					)...,
				),
				"options method": baseOpts.PATCH(
					reqb.Path("/users/%d", 123),
					reqb.BodyBytes([]byte("{\"name\":\"bar\"}")),
					reqb.Header("Content-Type", "application/json"),
					reqb.Cookie(&http.Cookie{Name: "session", Value: "session1"}),
				),
			},
			&request{
				Method: http.MethodPatch,
				Url:    "/users/123",
				Header: expectedHeader(http.Header{
					"Content-Type": []string{"application/json"},
					"Cookie":       []string{"session=session1"},
				}),
				Body: "{\"name\":\"bar\"}",
			},
		},
		{
			"DELETE /users/456",
			map[string]reqb.Func{
				"ad-hoc": reqb.DELETE(
					reqb.BaseUrl(baseURL),
					reqb.Path("/users/%d", 456),
					reqb.BodyString(""),
				),
				"package function": reqb.DELETE(
					append(baseOpts,
						reqb.Path("/users/%d", 456),
						reqb.BodyString(""),
					)...,
				),
				"options method": baseOpts.DELETE(
					reqb.Path("/users/%d", 456),
					reqb.BodyString(""),
				),
			},
			&request{
				Method: http.MethodDelete,
				Url:    "/users/456",
				Header: expectedHeader(http.Header{}),
				Body:   "",
			},
		},
		{
			"OPTIONS /",
			map[string]reqb.Func{
				"ad-hoc with schema, host, port-string": reqb.OPTIONS(
					// reqb.BaseUrl(baseURL),
					reqb.Scheme("http"),
					reqb.Host(testServerURL.Hostname()),
					reqb.PortString(testServerURL.Port()),
				),
				"ad-hoc with baseUrl": reqb.OPTIONS(reqb.BaseUrl(baseURL)),
				"package function":    reqb.OPTIONS(baseOpts...),
				"options method":      baseOpts.OPTIONS(),
			},
			&request{
				Method: http.MethodOptions,
				Url:    "/",
				Header: expectedHeader(http.Header{}),
				Body:   "",
			},
		},
	}

	for _, p := range patterns {
		for funcName, getter := range p.funcs {
			t.Run(p.name+" "+funcName, func(t *testing.T) {
				client := &http.Client{}

				req := getter(t)
				resp, err := client.Do(req)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					t.Fatalf("unexpected status code: %d", resp.StatusCode)
				}
				cookies := resp.Cookies()
				t.Logf("CLIENT %d cookies: %+v", len(cookies), cookies)
				// client.Jar.SetCookies(req.URL, resp.Cookies())

				respBody, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				t.Logf("CLIENT: respBody%s", string(respBody))

				var actual request
				if err := json.Unmarshal(respBody, &actual); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				assert.Equal(t, p.expected, &actual)
			})
		}
	}
}
