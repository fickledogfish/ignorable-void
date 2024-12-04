package dto

import (
	"time"

	"example.com/services/posts/core/domain"
)

type LatestPostsItemResponse struct {
	Id string `json:"id"`

	UserId    string    `json:"user_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

func NewLatestPostsItemResponseFromDomain(
	post domain.Post,
) LatestPostsItemResponse {
	return LatestPostsItemResponse{
		Id: post.Id.String(),

		UserId:    post.UserId.String(),
		Content:   post.Content,
		Timestamp: post.Timestamp,
	}
}
