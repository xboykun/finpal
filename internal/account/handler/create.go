package handler

import (
	"encoding/json"
	"net/http"
)

func CreateHandler(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode("Hallo")
}
