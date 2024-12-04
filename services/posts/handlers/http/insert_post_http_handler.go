package httphandler

import "net/http"

type insertPostHttpHandler interface {
	http.Handler
}
