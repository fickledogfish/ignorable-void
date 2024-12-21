package insertposthttpinboundadapter

import "example.com/services/posts/core/ports"

func New(
	service ports.InsertPostService,
) *insertPostHttpHandler {
	return &insertPostHttpHandler{
		service: service,
	}
}
