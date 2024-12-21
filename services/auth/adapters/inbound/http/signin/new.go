package signinhttpinboundadapter

import "example.com/services/auth/core/ports"

func New(
	service ports.SignInUserService,
) *signInHttpHandler {
	return &signInHttpHandler{
		service: service,
	}
}
