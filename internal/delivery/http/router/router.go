package router

import (
	"net/http"

	"nami/internal/delivery/http/handler"
	"nami/internal/delivery/http/lib"
	"nami/internal/entity"
)

func AccountRouter(h *http.ServeMux) {
	h.Handle("POST /api/v1/account/create", lib.
		HttpHandleChain(
			lib.HttpParseInput[entity.CreateAccountInput](
				lib.WithHttpErrorHandler(lib.DefaultMiddlewareErrorHandler),
			),
		).
		ThenFunc(handler.CreateAccount),
	)

	h.Handle("GET /api/v1/account/detail", lib.
		HttpHandleChain().
		ThenFunc(nil),
	)

	h.Handle("PATCH /api/v1/account/status", lib.
		HttpHandleChain().
		ThenFunc(nil),
	)
}
