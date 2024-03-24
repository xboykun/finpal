package lib

import (
	"net/http"

	"github.com/ggicci/httpin"
)

func HttpTakeInput[T any](r *http.Request) *T {
	return r.Context().Value(httpin.Input).(*T)
}
