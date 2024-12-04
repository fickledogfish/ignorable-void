package page

type Page[Content any] struct {
	Number     int
	IsLast     bool
	TotalPages int

	Content []Content
}

func New[Content any](
	number int,
	totalPages int,
	content []Content,
) Page[Content] {
	return Page[Content]{
		Number:     number,
		TotalPages: totalPages,
		IsLast:     number == totalPages,

		Content: content,
	}
}
