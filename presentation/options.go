package presentation

type Opts struct {
	URLParams   []string
	QueryParams []string
}

type OptsFunc func(*Opts)

var Default = DefaultOpts()

func With(options ...OptsFunc) Opts {
	opts := Opts{}

	for _, o := range options {
		o(&opts)
	}

	return opts
}

func DefaultOpts() Opts {
	return Opts{}
}

func UrlParams(params ...string) OptsFunc {
	return func(opts *Opts) {
		opts.URLParams = make([]string, len(params))

		for i := 0; i < len(params); i++ {
			opts.URLParams[i] = params[i]
		}
	}
}

func QueryParams(queryParams ...string) OptsFunc {
	return func(opts *Opts) {
		opts.QueryParams = make([]string, len(queryParams))

		for i := 0; i < len(queryParams); i++ {
			opts.QueryParams[i] = queryParams[i]
		}
	}
}
