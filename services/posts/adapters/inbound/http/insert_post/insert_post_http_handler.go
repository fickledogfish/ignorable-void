package insertposthttpinboundadapter

import (
	"encoding/json"
	"net/http"

	"example.com/core/responses"
	"example.com/core/uuid"
	"example.com/services/posts/core/domain"
	"example.com/services/posts/core/ports"
)

type insertPostHttpHandler struct {
	service ports.InsertPostService
}

func (h *insertPostHttpHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	var decodedBody insertPostRequestDto
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

	responses.Ok(w, r, insertPostResponseDto{
		Id: postId.String(),
	})
}
