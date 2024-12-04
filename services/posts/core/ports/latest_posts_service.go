package ports

import "example.com/services/posts/core/domain"

type LatestPostsService interface {
	LatestPosts(int) ([]domain.Post, error)
}
