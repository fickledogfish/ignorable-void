package httphandler

import "net/http"

type signUpHttpHandler interface {
	http.Handler
}
