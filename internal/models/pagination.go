package models

type Pagination[T any] struct {
	Items      []T   `json:"items"`
	TotalItems int64 `json:"totalItems"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"totalPages"`
}
