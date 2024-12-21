package main

import "net/http"

type latestPostsHttpHandler interface {
	http.Handler
}
