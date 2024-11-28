package reqb

import (
	"net/http"
	"testing"
)

// Func is a type alias for a function that takes a *testing.T and returns an *http.Request.
type Func = func(*testing.T) *http.Request
