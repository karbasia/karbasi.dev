package pagination

import (
	"math"
	"net/http"
	"strconv"

	"github.com/karbasia/karbasi.dev/internal/store"
)

func FromRequest(r *http.Request) store.PaginationParams {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 1000 {
		pageSize = 1000
	}

	return store.PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
}

func Offset(p store.PaginationParams) int {
	return (p.Page - 1) * p.PageSize
}

func TotalPages(totalItems, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	return int(math.Ceil(float64(totalItems) / float64(pageSize)))
}
