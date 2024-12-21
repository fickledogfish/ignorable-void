package main

import "net/http"

type insertPostHttpHandler interface {
	http.Handler
}
