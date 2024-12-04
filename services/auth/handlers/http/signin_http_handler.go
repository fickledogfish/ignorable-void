package httphandler

import "net/http"

type signInHttpHandler interface {
	http.Handler
}
