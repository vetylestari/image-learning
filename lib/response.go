package lib

import (
	"encoding/json"
	"net/http"
)

// respondwithJSON write json response format
func ResponseJSON(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
