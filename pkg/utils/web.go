package utils

import (
	"encoding/json"
	"net/http"
)

func SendRespond(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(data)
}
