package commom

import (
	"encoding/json"
	"net/http"
)

type CommonsHandler struct{}

func (c CommonsHandler) SendJson(w http.ResponseWriter, body map[string]interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}
