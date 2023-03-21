package core

import (
	"errors"
)

type TransformPresenter[TDataIn any, TDataOut any] struct {
	Result    TDataIn
	Err       error
	Converter func(TDataIn) (TDataOut, error) //IConverting[TDataIn, TDataOut]
}

func ConvertData[TDataIn any, TDataOut any](converter func(TDataIn) (TDataOut, error), data TDataIn) (TDataOut, error) {

	if converter != nil {
		return converter(data)
	}

	var zeroValue TDataOut
	return zeroValue, errors.New(CONVERTER_NOT_INJECT_MESSAGE)
}

func (this *TransformPresenter[TDataIn, TDataOut]) Present(data TDataIn, err error) {
	this.Result = data
	this.Err = err
}

func (this TransformPresenter[TDataIn, TDataOut]) Print() (TDataOut, error) {

	if this.Err != nil {
		var zeroval TDataOut
		return zeroval, this.Err
	}

	return ConvertData(this.Converter, this.Result)
}

func (this *TransformPresenter[TDataIn, TDataOut]) ZeroValueErrorTransformPresenter(err error) {
	var zeroVal TDataIn
	this.Present(zeroVal, err)
}
