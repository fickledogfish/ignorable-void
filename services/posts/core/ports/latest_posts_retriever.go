package ports

import "example.com/services/posts/core/domain"

type LatestPostsRetriever interface {
	GetLatestPosts(int) ([]domain.Post, error)
}
