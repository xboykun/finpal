package router

import (
	"net/http"

	"nami/internal/entity"
	"nami/internal/handler"
	"nami/pkg/api"
)

type Router struct {
	http.Handler
}

func NewServerRouter(mux http.Handler) *Router {
	return &Router{mux}
}

func (h Router) AccountRouter(h *http.ServeMux) {
	h.Handle("POST /api/v1/account/create", api.
		HandleChain(
			api.ParseInput[entity.CreateAccount](),
		).
		ThenFunc(handler.CreateAccount),
	)

	h.Handle("GET /api/v1/account/detail", api.
		HandleChain().
		ThenFunc(nil),
	)

	h.Handle("PATCH /api/v1/account/status", api.
		HandleChain().
		ThenFunc(nil),
	)
}
