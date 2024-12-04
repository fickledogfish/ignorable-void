package middlewares

import (
	"fmt"
	"net/http"
	"slices"

	"example.com/core/responses"
)

type allowedMethodsMiddleware struct {
	allowedMethods []string
	component      http.Handler
}

func AllowedMethods(
	methods []string,
	component http.Handler,
) allowedMethodsMiddleware {
	return allowedMethodsMiddleware{
		allowedMethods: methods,
		component:      component,
	}
}

func (m allowedMethodsMiddleware) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	if !slices.Contains(m.allowedMethods, r.Method) {
		w.WriteHeader(http.StatusBadRequest)
		responses.BadRequest(
			w,
			r,
			fmt.Sprintf("Unsupported method: %s", r.Method),
		)
		return
	}

	m.component.ServeHTTP(w, r)
}
