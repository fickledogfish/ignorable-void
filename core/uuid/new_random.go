package uuid

import guuid "github.com/google/uuid"

func NewRandom() (UUID, error) {
	if uuid, err := guuid.NewRandom(); err != nil {
		return Nil, err
	} else {
		return UUID(uuid), nil
	}
}
