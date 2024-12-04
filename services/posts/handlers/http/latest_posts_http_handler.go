package httphandler

import "net/http"

type latestPostsHttpHandler interface {
	http.Handler
}
