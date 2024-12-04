package httpadapters

import (
	"net/http"
	"strconv"

	"example.com/core/responses"
	"example.com/services/posts/adapters/http/dto"
	"example.com/services/posts/core/domain"
	"example.com/services/posts/core/ports"
)

type latestPostsHttpHandler struct {
	service ports.LatestPostsService
}

func NewLatestPostsHttpHandler(
	service ports.LatestPostsService,
) *latestPostsHttpHandler {
	return &latestPostsHttpHandler{
		service: service,
	}
}

func (h *latestPostsHttpHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	var count int
	if requestedCount, err := func() (int, error) {
		countParam := r.URL.Query().Get("max")
		if countParam == "" {
			return 0, nil
		}

		return strconv.Atoi(countParam)
	}(); err != nil {
		responses.InternalServerError(w, r, err.Error())
	} else {
		count = requestedCount
	}

	var posts []domain.Post
	if latestPosts, err := h.service.LatestPosts(count); err != nil {
		responses.InternalServerError(w, r, err.Error())
	} else {
		posts = latestPosts
	}

	var response []dto.LatestPostsItemResponse
	for _, post := range posts {
		response = append(
			response,
			dto.NewLatestPostsItemResponseFromDomain(post),
		)
	}

	responses.Ok(w, r, response)
}
