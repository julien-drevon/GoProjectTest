package core

import (
	"errors"
)

type AsyncTransformPresenter[TDataIn any, TDataOut any] struct {
	Result      TDataIn
	Err         error
	Converter   IAsyncConverting[TDataIn, TDataOut]
	TransResult TDataOut
	TransError  error
}

func ConvertDataAsync[TDataIn any, TDataOut any](converter IAsyncConverting[TDataIn, TDataOut], data TDataIn) (TDataOut, error) {

	if converter != nil {
		chanResult := make(chan TDataOut)
		chanErr := make(chan error)
		go converter.ConvertAsync(data, chanResult, chanErr)
		return <-chanResult, <-chanErr
	}

	var zeroValue TDataOut
	return zeroValue, errors.New(CONVERTER_NOT_INJECT_MESSAGE)
}

func (presenter *AsyncTransformPresenter[TDataIn, TDataOut]) Present(data TDataIn, err error) {
	presenter.Result = data
	presenter.Err = err
	presenter.TransResult, presenter.TransError = ConvertDataAsync(presenter.Converter, presenter.Result)
}

func (presenter AsyncTransformPresenter[TDataIn, TDataOut]) Print() (TDataOut, error) {

	if presenter.Err != nil {
		var zeroval TDataOut
		return zeroval, presenter.Err
	}

	return presenter.TransResult, presenter.TransError
}

func (this *AsyncTransformPresenter[TDataIn, TDataOut]) ZeroValueErrorTransformPresenter(err error) {
	var zeroVal TDataIn
	this.Present(zeroVal, err)
}
