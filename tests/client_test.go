package testrequest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/akm/go-testrequest"
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
	factory := testrequest.NewFactory(testrequest.BaseUrl(baseURL))

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
		name         string
		reqByFunc    testrequest.Func
		reqByFactory testrequest.Func
		expected     *request
	}
	patterns := []pattern{
		{
			"GET /",
			testrequest.GET(testrequest.BaseUrl(baseURL)),
			factory.GET(),
			&request{
				Method: http.MethodGet,
				Url:    "/",
				Header: expectedHeader(http.Header{}),
				Body:   "",
			},
		},
		{
			"POST /users",
			testrequest.POST(
				testrequest.BaseUrl(baseURL),
				testrequest.Path("/users"),
				testrequest.BodyString("hello, world"),
			),
			factory.POST(
				testrequest.Path("/users"),
				testrequest.BodyString("hello, world"),
			),
			&request{
				Method: http.MethodPost,
				Url:    "/users",
				Header: expectedHeader(http.Header{}),
				Body:   "hello, world",
			},
		},
		{
			"PUT /users/123",
			testrequest.PUT(
				testrequest.BaseUrl(baseURL),
				testrequest.Path("/users/%d", 123),
				testrequest.BodyString("{\"name\":\"foo\"}"),
				testrequest.Header("Content-Type", "application/json"),
			),
			factory.PUT(
				testrequest.Path("/users/%d", 123),
				testrequest.BodyString("{\"name\":\"foo\"}"),
				testrequest.Header("Content-Type", "application/json"),
			),
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
			testrequest.PATCH(
				testrequest.BaseUrl(baseURL),
				testrequest.Path("/users/%d", 123),
				testrequest.BodyBytes([]byte("{\"name\":\"bar\"}")),
				testrequest.Header("Content-Type", "application/json"),
				testrequest.Cookie(&http.Cookie{Name: "session", Value: "session1"}),
			),
			factory.PATCH(
				testrequest.Path("/users/%d", 123),
				testrequest.BodyBytes([]byte("{\"name\":\"bar\"}")),
				testrequest.Header("Content-Type", "application/json"),
				testrequest.Cookie(&http.Cookie{Name: "session", Value: "session1"}),
			),
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
			testrequest.DELETE(
				testrequest.BaseUrl(baseURL),
				testrequest.Path("/users/%d", 456),
				testrequest.BodyString(""),
			),
			factory.DELETE(
				testrequest.Path("/users/%d", 456),
				testrequest.BodyString(""),
			),
			&request{
				Method: http.MethodDelete,
				Url:    "/users/456",
				Header: expectedHeader(http.Header{}),
				Body:   "",
			},
		},
		{
			"OPTIONS /",
			testrequest.OPTIONS(
				// testrequest.BaseUrl(baseURL),
				testrequest.Scheme("http"),
				testrequest.Host(testServerURL.Hostname()),
				testrequest.PortString(testServerURL.Port()),
			),
			factory.OPTIONS(),
			&request{
				Method: http.MethodOptions,
				Url:    "/",
				Header: expectedHeader(http.Header{}),
				Body:   "",
			},
		},
	}

	type getter struct {
		name string
		get  func(p pattern) testrequest.Func
	}

	funcGetters := []getter{
		{"by func", func(p pattern) testrequest.Func { return p.reqByFunc }},
		{"by factory", func(p pattern) testrequest.Func { return p.reqByFactory }},
	}

	for _, p := range patterns {
		for _, fg := range funcGetters {
			getter := fg.get(p)
			t.Run(p.name+" "+fg.name, func(t *testing.T) {
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
