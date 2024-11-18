package testrequest

func New(method string, options ...Option) Func {
	return DefaultFactory.New(method, options...)
}

func GET(v ...Option) Func     { return DefaultFactory.GET(v...) }
func HEAD(v ...Option) Func    { return DefaultFactory.HEAD(v...) }
func POST(v ...Option) Func    { return DefaultFactory.POST(v...) }
func PUT(v ...Option) Func     { return DefaultFactory.PUT(v...) }
func PATCH(v ...Option) Func   { return DefaultFactory.PATCH(v...) }
func DELETE(v ...Option) Func  { return DefaultFactory.DELETE(v...) }
func CONNECT(v ...Option) Func { return DefaultFactory.CONNECT(v...) }
func OPTIONS(v ...Option) Func { return DefaultFactory.OPTIONS(v...) }
func TRACE(v ...Option) Func   { return DefaultFactory.TRACE(v...) }
