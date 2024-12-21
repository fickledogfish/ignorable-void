package main

import "net/http"

type signUpHttpHandler interface {
	http.Handler
}
