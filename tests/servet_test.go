package testrequest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akm/go-testrequest"
	"github.com/stretchr/testify/assert"
)

func TestWithServer(t *testing.T) {
	testServer := startEchoServer()
	testServer.Start()
	defer testServer.Close()

	type pattern *struct {
		req      *http.Request
		expected *request
	}
	patterns := []pattern{
		{
			req: testrequest.GET(),
			expected: &request{
				Method: http.MethodGet,
				Url:    "/",
				Header: http.Header{},
				Body:   "",
			},
		},
		{
			req: testrequest.POST(
				testrequest.Path("/users"),
				testrequest.BodyString("hello, world"),
			),
			expected: &request{
				Method: http.MethodPost,
				Url:    "/users",
				Header: http.Header{},
				Body:   "hello, world",
			},
		},
		{
			req: testrequest.PUT(
				testrequest.Path("/users/%d", 123),
				testrequest.BodyString("{\"name\":\"foo\"}"),
				testrequest.Header("Content-Type", "application/json"),
			),
			expected: &request{
				Method: http.MethodPut,
				Url:    "/users/123",
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body: "{\"name\":\"foo\"}",
			},
		},
		{
			req: testrequest.PATCH(
				testrequest.Path("/users/%d", 123),
				testrequest.BodyBytes([]byte("{\"name\":\"bar\"}")),
				testrequest.Header("Content-Type", "application/json"),
				testrequest.Cookie(&http.Cookie{Name: "session", Value: "123"}),
			),
			expected: &request{
				Method: http.MethodPatch,
				Url:    "/users/123",
				Header: http.Header{
					"Content-Type": []string{"application/json"},
					"Cookie":       []string{"session=123"},
				},
				Body: "{\"name\":\"bar\"}",
			},
		},
		{
			req: testrequest.DELETE(
				testrequest.Path("/users/%d", 456),
				testrequest.BodyString(""),
			),
			expected: &request{
				Method: http.MethodDelete,
				Url:    "/users/456",
				Body:   "",
			},
		},
	}

	for _, p := range patterns {
		t.Run(fmt.Sprintf("%s %s", p.req.Method, p.req.URL.Path), func(t *testing.T) {
			resp, err := http.DefaultClient.Do(p.req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("unexpected status code: %d", resp.StatusCode)
			}

			var actual request
			if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			assert.Equal(t, p.expected, &actual)
		})
	}

}

type request struct {
	Method string      `json:"method"`
	Url    string      `json:"url"`
	Header http.Header `json:"header"`
	Body   string      `json:"body"`
}

func startEchoServer() *httptest.Server {
	echoHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(request{
			Method: r.Method,
			Url:    r.URL.String(),
			Header: r.Header,
			Body:   string(body),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
	return httptest.NewUnstartedServer(echoHandler)
}
