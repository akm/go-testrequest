package reqb

type Options []Option

func (o Options) Append(opts ...Option) Options {
	return append(o, opts...)
}

func (o Options) With(opts ...Option) Options {
	return o.Append(opts...)
}

func (o Options) GET(v ...Option) Func     { return GET(o.With(v...)...) }
func (o Options) HEAD(v ...Option) Func    { return HEAD(o.With(v...)...) }
func (o Options) POST(v ...Option) Func    { return POST(o.With(v...)...) }
func (o Options) PUT(v ...Option) Func     { return PUT(o.With(v...)...) }
func (o Options) PATCH(v ...Option) Func   { return PATCH(o.With(v...)...) }
func (o Options) DELETE(v ...Option) Func  { return DELETE(o.With(v...)...) }
func (o Options) CONNECT(v ...Option) Func { return CONNECT(o.With(v...)...) }
func (o Options) OPTIONS(v ...Option) Func { return OPTIONS(o.With(v...)...) }
func (o Options) TRACE(v ...Option) Func   { return TRACE(o.With(v...)...) }
