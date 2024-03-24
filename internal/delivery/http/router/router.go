package router

import (
	"net/http"

	"nami/internal/delivery/http/handler"
	"nami/internal/delivery/http/middleware"
	"nami/internal/entity"
)

func AccountRouter(h *http.ServeMux) {
	h.Handle("POST /api/v1/account/create", middleware.
		Chain(
			middleware.NewInput[entity.CreateAccountInput](
				middleware.InputParseErrorHandler,
			),
		).
		ThenFunc(handler.CreateAccount),
	)

	h.Handle("POST /api/v1/account/detail", middleware.
		Chain().
		ThenFunc(nil),
	)

	h.Handle("POST /api/v1/account/status", middleware.
		Chain().
		ThenFunc(nil),
	)
}
