package page

type Dto[Content any] struct {
	Number     int  `json:"number"`
	IsLast     bool `json:"is_last"`
	TotalPages int  `json:"total_pages"`
	Count      int  `json:"count"`

	Content []Content `json:"content"`
}

func NewDto[Content any](
	number int,
	totalPages int,
	content []Content,
) Dto[Content] {
	return Dto[Content]{
		Number:     number,
		TotalPages: totalPages,
		IsLast:     number == totalPages,
		Count:      len(content),

		Content: content,
	}
}

func NewDtoFromPage[PageContent any, DtoContent any](
	page Page[PageContent],
	mapper func(PageContent) DtoContent,
) Dto[DtoContent] {
	content := make([]DtoContent, len(page.Content))
	for index, element := range page.Content {
		content[index] = mapper(element)
	}

	return NewDto(
		page.Number,
		page.TotalPages,
		content,
	)
}
