package main

import (
	"fmt"
	"net/http"

	"example.com/core/middlewares"
)

type httpHandler struct {
	port          string
	signInHandler signInHttpHandler
	signUpHandler signUpHttpHandler
}

func newHttpHandler(
	port string,
	signInHandler signInHttpHandler,
	signUpHandler signUpHttpHandler,
) httpHandler {
	return httpHandler{
		port:          port,
		signInHandler: signInHandler,
		signUpHandler: signUpHandler,
	}
}

func (handler httpHandler) ListenAndServe() error {
	mux := http.NewServeMux()
	mux.Handle(
		"/signup",
		middlewares.ContentTypeHeaderWriter(
			middlewares.AllowedMethods(
				[]string{http.MethodPost},
				handler.signUpHandler,
			),
		),
	)
	mux.Handle(
		"/signin",
		middlewares.ContentTypeHeaderWriter(
			middlewares.AllowedMethods(
				[]string{http.MethodPost},
				handler.signInHandler,
			),
		),
	)

	s := http.Server{
		Addr:    fmt.Sprintf(":%s", handler.port),
		Handler: mux,
	}

	return s.ListenAndServe()
}
