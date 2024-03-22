package handler

import (
	"net/http"

	"nami/internal/delivery/http/lib"
	"nami/internal/entity"
)

func CreateAccount(rw http.ResponseWriter, r *http.Request) {
	d := lib.HttpTakeInput[entity.CreateAccountInput](r)

	lib.NewHttpRes[*entity.AccountBodyJson]().
		SetData(d.Payload).
		Success().
		JSON(rw)
}
