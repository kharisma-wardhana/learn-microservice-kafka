package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, map[string]string{"error": message})
}

func GetMuxVars(req *http.Request) map[string]string {
	vars := mux.Vars(req)
	if vars == nil {
		return map[string]string{}
	}
	return vars
}
