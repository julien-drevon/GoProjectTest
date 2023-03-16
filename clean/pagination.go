package core

import (
	"math"
)

type PaginationResult[TData any] struct {
	Result     []TData
	Pagination Pagination
}

type Pagination struct {
	CurrentPage int  `json:"CurrentPage"`
	HasNext     bool `json:"hasNext"`
	HasPrevious bool `json:"hasPrevious"`
	TotalPage   int  `json:"totalPage"`
	PageSize    int  `json:"pageSize"`
	Total       int  `json:"total"`
}

func NewPaginationResult[TData any](values []TData, total int, currentPage int, pageSizeRequest int) PaginationResult[TData] {
	pageNumber := SetToMin1(currentPage)
	totalElement := SetToMin0(total)
	totalPage := ComputeNbOfPage(totalElement, len(values))
	pageSizeDesired := SetToMin0(pageSizeRequest)

	return PaginationResult[TData]{
		Result: values,
		Pagination: Pagination{
			Total:       totalElement,
			TotalPage:   totalPage,
			CurrentPage: pageNumber,
			HasNext:     pageNumber < totalPage,
			HasPrevious: pageNumber > 1,
			PageSize:    pageSizeDesired,
		}}
}

func SetToMin0(currentPage int) int {
	if currentPage < 1 {
		return 0
	}

	return currentPage
}
func SetToMin1(currentPage int) int {
	if currentPage < 1 {
		return 1
	}

	return currentPage
}
func ComputeNbOfPage(totalEntry int, nbElementPerPage int) int {
	if nbElementPerPage == 0 {
		return 0
	}

	return int(math.Ceil(float64(float64(totalEntry) / float64(nbElementPerPage))))
}
