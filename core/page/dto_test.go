package page_test

import (
	"testing"

	"example.com/core/page"
	"github.com/stretchr/testify/require"
)

// TestNewDtoFromPage calls NewDtoFromPage, checking if the transformation was
// applied and the fields were set correctly.
func TestNewDtoFromPage(t *testing.T) {
	pageNumber := 1
	content := []int{1, 2, 3, 4}
	transformedContent := []float32{2, 4, 6, 8}

	original := page.New(pageNumber, pageNumber, content)

	dto := page.NewDtoFromPage(original, func(element int) float32 {
		return float32(element) * 2
	})

	require.Equal(t, pageNumber, dto.Number)
	require.Equal(t, pageNumber, dto.TotalPages)
	require.True(t, dto.IsLast)
	require.InDeltaSlice(t, transformedContent, dto.Content, 0.001)
}
