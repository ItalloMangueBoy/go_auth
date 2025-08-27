package views

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if status == http.StatusNoContent {
		return
	}

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(500)
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	}
}

// Message sends a JSON response with a message
func Message(w http.ResponseWriter, status int, err string) {
	JSON(w, status, map[string]string{"msg": err})
}
