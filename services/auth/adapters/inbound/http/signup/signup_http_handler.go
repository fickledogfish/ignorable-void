package signuphttpinboundadapter

import (
	"encoding/json"
	"net/http"

	"example.com/core/responses"
	"example.com/core/uuid"

	"example.com/services/auth/core/domain"
	"example.com/services/auth/core/ports"
)

type signUpHttpHandler struct {
	service ports.SignUpService
}

func (handler *signUpHttpHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	var decodedBody signUpRequestDto
	if err := json.NewDecoder(r.Body).Decode(&decodedBody); err != nil {
		responses.BadRequest(w, r, err.Error())
		return
	}

	var userId uuid.UUID
	if id, err := handler.service.SignUp(
		domain.SignUpRequest{
			Username: decodedBody.Username,
			Password: decodedBody.Password,
		},
	); err != nil {
		responses.InternalServerError(w, r, err.Error())
		return
	} else {
		userId = id
	}

	responses.Ok(w, r, signUpResponseDto{
		Id:       userId.String(),
		Username: decodedBody.Username,
	})
}
