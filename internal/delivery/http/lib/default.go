package lib

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ggicci/httpin"
	"github.com/ggicci/httpin/core"
	"github.com/justinas/alice"
)

const (
	_Created       = http.StatusCreated
	_Success       = http.StatusOK
	_Error         = http.StatusBadRequest
	_InternalError = http.StatusInternalServerError

	_DecodeErrorMessage  = "invalid decode request"
	_DecodeErrorTemplate = `{"code": %s,"message":%s, "data": null}`
	_EncodeErrorMessage  = "invalid encode response"
)

var (
	_CreatedMessage       string = http.StatusText(http.StatusCreated)
	_SuccessMessage       string = http.StatusText(http.StatusOK)
	_ErrorMessage         string = http.StatusText(http.StatusBadRequest)
	_InternalErrorMessage string = http.StatusText(http.StatusInternalServerError)
)

var (
	HttpHandleChain = alice.New

	DefaultInternalErrorHandler = func(rw http.ResponseWriter, code, msg *string) {
		if len(*code) <= 0 {
			*code = strconv.FormatInt(int64(_InternalError), 10)
		}

		if len(*msg) <= 0 {
			*msg = _EncodeErrorMessage
		}

		rw.Write([]byte(fmt.Sprintf(_DecodeErrorTemplate, *code, *msg)))
	}

	DefaultMiddlewareErrorHandler = func(rw http.ResponseWriter, r *http.Request, err error) {
		NewHttpRes[any]().SetMessage(_DecodeErrorMessage).Error().JSON(rw)
	}
)

type (
	Option core.Option
)

func WithHttpErrorHandler(custom func(w http.ResponseWriter, r *http.Request, err error)) Option {
	return Option(core.WithErrorHandler(custom))
}

func WithHttpMaxMemory(maxMemory int64) Option {
	return Option(core.WithMaxMemory(maxMemory))
}

func WithHttpNestedDirectivesEnabled(enable bool) Option {
	return Option(core.WithNestedDirectivesEnabled(enable))
}

func HttpParseInput[T any](opts ...Option) func(http.Handler) http.Handler {
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
