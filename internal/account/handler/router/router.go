package router

import (
	"net/http"

	"github.com/justinas/alice"
)

func AccountRouter(h *http.ServeMux) {
	h.Handle("POST /api/v1/account/create", alice.New().Then(nil))
	h.Handle("GET /api/v1/account/detail", alice.New().Then(nil))
	h.Handle("PATCH /api/v1/account/status", alice.New().Then(nil))
}
