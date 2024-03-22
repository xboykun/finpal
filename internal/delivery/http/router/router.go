package router

import (
	"net/http"

	"nami/internal/delivery/http/handler"
	"nami/internal/delivery/http/lib"
	"nami/internal/entity"
)

func AccountRouter(h *http.ServeMux) {
	h.Handle("POST /api/v1/account/create", lib.
		HttpMiddlewareChain(
			lib.HttpInputParser[entity.CreateAccountInput](lib.WithHttpErrorParser(lib.DefaultMiddlewareError)),
		).
		ThenFunc(handler.CreateAccount),
	)

	h.Handle("GET /api/v1/account/detail", lib.
		HttpMiddlewareChain().
		ThenFunc(nil),
	)

	h.Handle("PATCH /api/v1/account/status", lib.
		HttpMiddlewareChain().
		ThenFunc(nil),
	)
}
