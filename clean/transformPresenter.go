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

func (presenter *TransformPresenter[TDataIn, TDataOut]) Present(data TDataIn, err error) {
	presenter.Result = data
	presenter.Err = err
}

func (presenter TransformPresenter[TDataIn, TDataOut]) Print() (TDataOut, error) {

	if presenter.Err != nil {
		var zeroval TDataOut
		return zeroval, presenter.Err
	}

	return ConvertData(presenter.Converter, presenter.Result)
}

func (this *TransformPresenter[TDataIn, TDataOut]) ZeroValueErrorTransformePresenter(err error) {
	var zeroVal TDataIn
	this.Present(zeroVal, err)
}
