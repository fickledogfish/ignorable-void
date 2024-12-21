package latestpostshttpinboundadapter

import (
	"time"

	"example.com/services/posts/core/domain"
)

type latestPostsItemResponse struct {
	Id string `json:"id"`

	UserId    string    `json:"user_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

func newLatestPostsItemResponseFromDomain(
	post domain.Post,
) latestPostsItemResponse {
	return latestPostsItemResponse{
		Id: post.Id.String(),

		UserId:    post.UserId.String(),
		Content:   post.Content,
		Timestamp: post.Timestamp,
	}
}
