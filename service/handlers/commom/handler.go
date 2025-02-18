package commom

import (
	"encoding/json"
	"net/http"
)

type CommonsHandler struct{}

func (c CommonsHandler) SendJson(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
