package views

import (
	"auth/src/helpers"
	"net/http"
)

// ModelErrors sends a JSON response with validation errors for a struct
func ModelErrors(w http.ResponseWriter, status int, msgs helpers.ErrorMessages, description ...string) {
	res := map[string]interface{}{
		"errors": msgs,
	}

	if len(description) > 0 {
		res["description"] = description[0]
	}

	JSON(w, status, res)
}
