package handler

import (
	"nami/internal/entity"
	"nami/pkg/api"
	"net/http"
)

func CreateAccount(rw http.ResponseWriter, r *http.Request) {
	input := api.TakeInput[entity.CreateAccount](r)

	api.NewResponse[*entity.CreateAccount]().
		SetData(input).
		Success().
		JSON(rw)
}
