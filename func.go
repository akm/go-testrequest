package testrequest

import (
	"net/http"
	"testing"
)

type Func = func(*testing.T) *http.Request
