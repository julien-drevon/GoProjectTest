package core

type PaginationResult[TData any] struct {
	Result []TData
	Pagination
}
