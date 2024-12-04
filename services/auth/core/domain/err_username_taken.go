package domain

type ErrUsernameTaken struct {
}

func NewErrUsernameTaken() ErrUsernameTaken {
	return ErrUsernameTaken{}
}

func (e ErrUsernameTaken) Error() string {
	return "username already taken"
}
