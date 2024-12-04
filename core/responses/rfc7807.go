package responses

import "net/http"

type RFC7807 struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func NewRFC7807(
	kind string,
	title string,
	status int,
	detail string,
	instance string,
) RFC7807 {
	return RFC7807{
		Type:     kind,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
	}
}

func NewRFC7807ForRequest(
	r *http.Request,
	title string,
	status int,
	detail string,
) RFC7807 {
	return RFC7807{
		Type:     r.Host,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: r.Host,
	}
}
