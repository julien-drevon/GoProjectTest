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
