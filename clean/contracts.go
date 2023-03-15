package core

import (
	"time"
)

type IPresentIn[TData any] interface {
	Present(TData, error)
}

type IPresentOut[TData any] interface {
	Print() (TData, error)
}

type IExecuteUseCase[TQuery any, TData any] interface {
	Execute(TQuery, IPresentIn[TData])
}

type IProvideDateTime interface {
	GetDateTime() (time.Time, error)
}

type IConverting[TDataIn any, TDataOut any] interface {
	Convert(TDataIn) (TDataOut, error)
}

type Pagination struct {
	CurrentPage int   `json:"CurrentPage"`
	HasNext     bool  `json:"hasNext"`
	HasPrevious bool  `json:"hasPrevious"`
	IsFirst     bool  `json:"isFirst"`
	IsLast      bool  `json:"isLast"`
	PageNumber  int64 `json:"pageNumber"`
	PageIndex   int   `json:"pageIndex"`
	PageSize    int   `json:"pageSize"`
	Total       int   `json:"total"`
}
