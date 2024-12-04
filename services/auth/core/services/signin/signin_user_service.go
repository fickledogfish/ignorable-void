package signinserv

import (
	"errors"

	domain "example.com/services/auth/core/domain"
	ports "example.com/services/auth/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type signInUserService struct {
	userIdRetriever ports.UserRetriever
}

func New(
	userIdRetriever ports.UserRetriever,
) *signInUserService {
	return &signInUserService{
		userIdRetriever: userIdRetriever,
	}
}

func (service *signInUserService) WithCredentials(
	credentials domain.UserSignInCredentials,
) (domain.AccessToken, error) {
	var user domain.User
	if retrievedUser, err := service.userIdRetriever.RetrieveUserWithName(
		credentials.Username,
	); err != nil {
		return domain.AccessToken{}, err
	} else if retrievedUser == nil {
		return domain.AccessToken{}, errors.New("username not found")
	} else {
		user = *retrievedUser
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password),
	); err != nil {
		return domain.AccessToken{}, err
	}

	return domain.AccessToken{
		Id:          user.Id,
		AccessToken: "",
	}, nil
}
