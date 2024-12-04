package latestpostsserv

import (
	"example.com/services/posts/core/domain"
	"example.com/services/posts/core/ports"
)

type latestPostsService struct {
	retriever ports.LatestPostsRetriever
}

func New(
	retriever ports.LatestPostsRetriever,
) *latestPostsService {
	return &latestPostsService{
		retriever: retriever,
	}
}

func (service *latestPostsService) LatestPosts(
	count int,
) ([]domain.Post, error) {
	return service.retriever.GetLatestPosts(min(max(count, 1), 100))
}
