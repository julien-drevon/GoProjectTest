package core

import (
	"errors"
)

type TransformPresenter[TDataIn any, TDataOut any] struct {
	Result    TDataIn
	Err       error
	Converter func(TDataIn, error) (TDataOut, error)
}

func ConvertData[TDataIn any, TDataOut any](converter func(TDataIn, error) (TDataOut, error), data TDataIn, err error) (TDataOut, error) {

	if converter != nil {
		return converter(data, err)
	}

	var zeroValue TDataOut
	return zeroValue, errors.New(CONVERTER_NOT_INJECT_MESSAGE)
}

func (this *TransformPresenter[TDataIn, TDataOut]) Present(data TDataIn, err error) {
	this.Result = data
	this.Err = err
}

func (this TransformPresenter[TDataIn, TDataOut]) Print() (TDataOut, error) {
	return ConvertData(this.Converter, this.Result, this.Err)
}

func (this *TransformPresenter[TDataIn, TDataOut]) ZeroValueErrorTransformPresenter(err error) {
	var zeroVal TDataIn
	this.Present(zeroVal, err)
}
