package store

type PaginationParams struct {
	Page     int
	PageSize int
}

type PaginationMeta struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type PaginatedResult[T any] struct {
	Items      []T
	Pagination PaginationMeta
}
