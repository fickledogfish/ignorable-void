package responses

import (
	"encoding/json"
	"net/http"
)

func Ok(
	w http.ResponseWriter,
	r *http.Request,
	data any,
) {
	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		InternalServerError(w, r, err.Error())
		return
	}
}
