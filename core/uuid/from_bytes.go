package uuid

import guuid "github.com/google/uuid"

func Parse(s string) (UUID, error) {
	if uuid, err := guuid.Parse(s); err != nil {
		return Nil, err
	} else {
		return UUID(uuid), nil
	}
}
