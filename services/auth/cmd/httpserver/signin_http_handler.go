package main

import "net/http"

type signInHttpHandler interface {
	http.Handler
}
