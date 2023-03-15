package core

import (
	"math"
)

type PaginationResult[TData any] struct {
	Result     []TData
	Pagination Pagination
}

type Pagination struct {
	//CurrentPage int  `json:"CurrentPage"`
	//HasNext     bool `json:"hasNext"`
	//HasPrevious bool `json:"hasPrevious"`
	//IsFirst     bool `json:"isFirst"`
	//IsLast      bool `json:"isLast"`
	TotalPage int `json:"totalPage"`
	//PageIndex   int  `json:"pageIndex"`
	//PageSize    int  `json:"pageSize"`
	Total int `json:"total"`
}

func NewPaginationResult[TData any](values []TData, total int) PaginationResult[TData] {
	totalPage := ComputeNbOfPage(total, len(values))
	return PaginationResult[TData]{
		Result: values,
		Pagination: Pagination{
			Total:     total,
			TotalPage: totalPage,
		}}
}

func ComputeNbOfPage(totalEntry int, nbElementPerPage int) int {
	if nbElementPerPage == 0 {
		return 1
	}

	return int(math.Ceil(float64(float64(totalEntry) / float64(nbElementPerPage))))
}
