package core

import (
	"time"
)

const CONVERTER_NOT_INJECT_MESSAGE = "Converter must be inject in TransformerPresenter"

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
	GetDateTime() time.Time
}

type IConverting[TDataIn any, TDataOut any] interface {
	Convert(TDataIn) (TDataOut, error)
}

type IAsyncConverting[TDataIn any, TDataOut any] interface {
	ConvertAsync(TDataIn, chan TDataOut, chan error)
}
