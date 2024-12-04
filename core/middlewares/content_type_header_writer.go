package middlewares

import (
	"net/http"
)

type contentTypeHeaderWriterMiddleware struct {
	component http.Handler
}

func ContentTypeHeaderWriter(
	component http.Handler,
) contentTypeHeaderWriterMiddleware {
	return contentTypeHeaderWriterMiddleware{
		component: component,
	}
}

func (m contentTypeHeaderWriterMiddleware) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	m.component.ServeHTTP(w, r)
}
