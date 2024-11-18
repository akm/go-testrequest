package testrequest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
)

type request struct {
	Method string      `json:"method"`
	Url    string      `json:"url"`
	Header http.Header `json:"header"`
	Body   string      `json:"body"`
}

func startEchoServer(logger interface {
	Logf(format string, args ...any)
}) *httptest.Server {
	echoHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("SERVER: %s %s", r.Method, r.URL.String())
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		cookie := r.Cookies()
		logger.Logf("SERVER: cookie: %+v", cookie)
		for _, c := range cookie {
			cookie := &http.Cookie{
				Name:  c.Name,
				Value: c.Value,
				Path:  "/",
			}
			logger.Logf("SERVER: set-cookie: %+v\n", *c)
			http.SetCookie(w, cookie)
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
