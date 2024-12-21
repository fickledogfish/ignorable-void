package insertpostserv

import (
	"example.com/core/uuid"
	"example.com/services/posts/core/domain"
	"example.com/services/posts/core/ports"
)

type insertPostService struct {
	newPostInserter ports.NewPostInserter
}

func New(
	newPostInserter ports.NewPostInserter,
) *insertPostService {
	return &insertPostService{
		newPostInserter: newPostInserter,
	}
}

func (service *insertPostService) Post(
	post domain.Post,
) (uuid.UUID, error) {
	if id, err := uuid.NewRandom(); err != nil {
		return uuid.Nil, err
	} else {
		post.Id = id
	}

	if err := service.newPostInserter.InsertNewPost(post); err != nil {
		return uuid.Nil, err
	}

	return post.Id, nil
}
