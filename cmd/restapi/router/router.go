package router

import (
	accountRouter "nami/internal/handler/router"
	"net/http"
)

func Init(h *http.ServeMux) {
	accountRouter.AccountRouter(h)
}
