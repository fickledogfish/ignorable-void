package responses

import (
	"encoding/json"
	"net/http"
)

func BadRequest(
	w http.ResponseWriter,
	r *http.Request,
	details string,
) {
	json.NewEncoder(w).Encode(NewRFC7807ForRequest(
		r,
		"Bad request",
		http.StatusBadRequest,
		details,
	))
}
