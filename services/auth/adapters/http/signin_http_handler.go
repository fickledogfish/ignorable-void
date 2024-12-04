package httpadapters

import (
	"encoding/json"
	"net/http"

	"example.com/core/responses"
	"example.com/services/auth/adapters/http/dto"
	"example.com/services/auth/core/domain"
	"example.com/services/auth/core/ports"
)

type signInHttpHandler struct {
	service ports.SignInUserService
}

func NewSignInHttpHandler(
	service ports.SignInUserService,
) *signInHttpHandler {
	return &signInHttpHandler{
		service: service,
	}
}

func (handler *signInHttpHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	var decodedBody dto.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&decodedBody); err != nil {
		responses.BadRequest(w, r, err.Error())
		return
	}

	var accessToken domain.AccessToken
	if token, err := handler.service.WithCredentials(
		domain.UserSignInCredentials{
			Username: decodedBody.Username,
			Password: decodedBody.Password,
		},
	); err != nil {
		responses.InternalServerError(w, r, err.Error())
		return
	} else {
		accessToken = token
	}

	responses.Ok(w, r, dto.SignInResponse{
		Id:          accessToken.Id.String(),
		AccessToken: accessToken.AccessToken,
	})
}
