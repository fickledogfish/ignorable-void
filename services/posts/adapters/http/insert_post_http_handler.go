package httpadapters

import (
	"encoding/json"
	"net/http"

	"example.com/core/responses"
	"example.com/services/posts/adapters/http/dto"
	"example.com/services/posts/core/domain"
	"example.com/services/posts/core/ports"
	"github.com/google/uuid"
)

type insertPostHttpHandler struct {
	service ports.InsertPostService
}

func NewInsertPostHttpHandler(
	service ports.InsertPostService,
) *insertPostHttpHandler {
	return &insertPostHttpHandler{
		service: service,
	}
}

func (h *insertPostHttpHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	var decodedBody dto.InsertPostRequest
	if err := json.NewDecoder(r.Body).Decode(&decodedBody); err != nil {
		responses.BadRequest(w, r, err.Error())
		return
	}

	var postId uuid.UUID
	if id, err := h.service.Post(domain.NewFreshPost(
		decodedBody.Content,
	)); err != nil {
		responses.InternalServerError(w, r, err.Error())
		return
	} else {
		postId = id
	}

	responses.Ok(w, r, dto.InsertPostResponse{
		Id: postId.String(),
	})
}
