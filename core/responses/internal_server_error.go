package responses

import (
	"encoding/json"
	"net/http"
)

func InternalServerError(
	w http.ResponseWriter,
	r *http.Request,
	details string,
) {
	json.NewEncoder(w).Encode(NewRFC7807ForRequest(
		r,
		"Internal server error",
		http.StatusBadRequest,
		details,
	))
}
