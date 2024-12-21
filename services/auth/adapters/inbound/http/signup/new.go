package signuphttpinboundadapter

import "example.com/services/auth/core/ports"

func New(
	service ports.SignUpService,
) *signUpHttpHandler {
	return &signUpHttpHandler{
		service: service,
	}
}
