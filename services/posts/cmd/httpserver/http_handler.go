package main

import (
	"fmt"
	"net/http"

	"example.com/core/middlewares"
)

type httpHandler struct {
	port string

	insertPost  insertPostHttpHandler
	latestPosts latestPostsHttpHandler
}

func newHttpHandler(
	port string,
	insertPost insertPostHttpHandler,
	latestPosts latestPostsHttpHandler,
) httpHandler {
	return httpHandler{
		port: port,

		insertPost:  insertPost,
		latestPosts: latestPosts,
	}
}

func (handler httpHandler) ListenAndServe() error {
	mux := http.NewServeMux()

	mux.Handle("/latest", middlewares.AllowedMethods(
		[]string{http.MethodGet},
		handler.latestPosts,
	))

	mux.Handle("/post", middlewares.AllowedMethods(
		[]string{http.MethodPost},
		handler.insertPost,
	))

	s := http.Server{
		Addr:    fmt.Sprintf(":%s", handler.port),
		Handler: mux,
	}

	return s.ListenAndServe()
}
