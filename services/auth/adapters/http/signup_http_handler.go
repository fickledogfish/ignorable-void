package httpadapters

import (
	"encoding/json"
	"net/http"

	"example.com/core/responses"
	"example.com/services/auth/adapters/http/dto"
	"example.com/services/auth/core/domain"
	"example.com/services/auth/core/ports"
	"github.com/google/uuid"
)

type signUpHttpHandler struct {
	service ports.SignUpService
}

func NewSignUpHttpHandler(
	service ports.SignUpService,
) *signUpHttpHandler {
	return &signUpHttpHandler{
		service: service,
	}
}

func (handler *signUpHttpHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	var decodedBody dto.SignInRequest
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

	responses.Ok(w, r, dto.SignUpResponse{
		Id:       userId.String(),
		Username: decodedBody.Username,
	})
}
