package middleware

import (
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/ggicci/httpin/core"

	"nami/internal/delivery/http/lib"
)

var (
	InputParseErrorHandler = WithErrorParser(lib.DefaultParseInputError)
)

type (
	Option core.Option
)

func WithErrorParser(custom func(w http.ResponseWriter, r *http.Request, err error)) Option {
	return Option(core.WithErrorHandler(custom))
}

func WithMaxMemory(maxMemory int64) Option {
	return Option(core.WithMaxMemory(maxMemory))
}

func WithNestedDirectivesEnabled(enable bool) Option {
	return Option(core.WithNestedDirectivesEnabled(enable))
}

func NewInput[T any](opts ...Option) func(http.Handler) http.Handler {
	var (
		d T
		o = make([]core.Option, len(opts))
	)

	for i, v := range opts {
		o[i] = core.Option(v)
	}

	return httpin.NewInput(d, o...)
}
