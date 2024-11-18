package testrequest

// New creates a new request with the specified method and options.
func New(method string, options ...Option) Func {
	return DefaultFactory.New(method, options...)
}

// GET creates a new GET request with the specified options.
func GET(v ...Option) Func { return DefaultFactory.GET(v...) }

// HEAD creates a new HEAD request with the specified options.
func HEAD(v ...Option) Func { return DefaultFactory.HEAD(v...) }

// POST creates a new POST request with the specified options.
func POST(v ...Option) Func { return DefaultFactory.POST(v...) }

// PUT creates a new PUT request with the specified options.
func PUT(v ...Option) Func { return DefaultFactory.PUT(v...) }

// PATCH creates a new PATCH request with the specified options.
func PATCH(v ...Option) Func { return DefaultFactory.PATCH(v...) }

// DELETE creates a new DELETE request with the specified options.
func DELETE(v ...Option) Func { return DefaultFactory.DELETE(v...) }

// CONNECT creates a new CONNECT request with the specified options.
func CONNECT(v ...Option) Func { return DefaultFactory.CONNECT(v...) }

// OPTIONS creates a new OPTIONS request with the specified options.
func OPTIONS(v ...Option) Func { return DefaultFactory.OPTIONS(v...) }

// TRACE creates a new TRACE request with the specified options.
func TRACE(v ...Option) Func { return DefaultFactory.TRACE(v...) }
