package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, resp interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("x-openrtb-version", "2.3")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}
