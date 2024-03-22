package lib

import (
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/ggicci/httpin/core"
)

type (
	Option core.Option
)

func WithHttpErrorParser(custom func(w http.ResponseWriter, r *http.Request, err error)) Option {
	return Option(core.WithErrorHandler(custom))
}

func WithHttpMaxMemory(maxMemory int64) Option {
	return Option(core.WithMaxMemory(maxMemory))
}

func WithHttpNestedDirectivesEnabled(enable bool) Option {
	return Option(core.WithNestedDirectivesEnabled(enable))
}

func HttpInputParser[T any](opts ...Option) func(http.Handler) http.Handler {
	var (
		d T
		o = make([]core.Option, len(opts))
	)

	for i, v := range opts {
		o[i] = core.Option(v)
	}

	return httpin.NewInput(d, o...)
}

func HttpTakeInput[T any](r *http.Request) *T {
	return r.Context().Value(httpin.Input).(*T)
}
