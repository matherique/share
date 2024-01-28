package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

func SendRespond(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(data)
}

var browersUserAgent = []string{
	"Mozilla",
	"Chrome",
}

func IsBrowerRequest(r *http.Request) bool {
	for _, b := range browersUserAgent {
		if strings.Contains(r.UserAgent(), b) {
			return true
		}
	}

	return false
}
