package signupserv

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"example.com/services/auth/core/domain"
	"example.com/services/auth/core/ports"
)

type signUpService struct {
	storer      ports.UserStorer
	idRetriever ports.UserRetriever
}

func New(
	storer ports.UserStorer,
	idRetriever ports.UserRetriever,
) *signUpService {
	return &signUpService{
		storer:      storer,
		idRetriever: idRetriever,
	}
}

func (service *signUpService) SignUp(
	credentials domain.SignUpRequest,
) (uuid.UUID, error) {
	if hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(credentials.Password),
		bcrypt.DefaultCost,
	); err != nil {
		return uuid.Nil, err
	} else {
		credentials.Password = string(hashedPassword)
	}

	var userId uuid.UUID
	if id, err := uuid.NewRandom(); err != nil {
		return uuid.Nil, err
	} else {
		userId = id
	}

	if retrievedUser, err := service.idRetriever.RetrieveUserWithName(
		credentials.Username,
	); err != nil {
		return uuid.Nil, err
	} else if retrievedUser != nil {
		return retrievedUser.Id, domain.NewErrUsernameTaken()
	}

	err := service.storer.StoreUser(userId, credentials)

	return userId, err
}
