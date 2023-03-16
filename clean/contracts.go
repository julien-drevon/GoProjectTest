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

// type IPresent[TIn any, TOut any] interface {
// 	IPresentIn[TIn]
// 	IPresentOut[TOut]
// }

type IExecuteUseCase[TQuery any, TData any] interface {
	Execute(TQuery, IPresentIn[TData])
}

type IProvideDateTime interface {
	GetDateTime() time.Time
}

type IConverting[TDataIn any, TDataOut any] interface {
	Convert(TDataIn) (TDataOut, error)
}
