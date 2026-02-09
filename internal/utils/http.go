package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter, status int, msg string) {
	SendJSON(w, status, map[string]string{"Error": msg})
}
