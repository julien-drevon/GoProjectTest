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

func (this *AsyncTransformPresenter[TDataIn, TDataOut]) Present(data TDataIn, err error) {
	this.Result = data
	this.Err = err
	this.TransResult, this.TransError = ConvertDataAsync(this.Converter, this.Result)
}

func (this AsyncTransformPresenter[TDataIn, TDataOut]) Print() (TDataOut, error) {

	if this.Err != nil {
		var zeroval TDataOut
		return zeroval, this.Err
	}

	return this.TransResult, this.TransError
}

func (this *AsyncTransformPresenter[TDataIn, TDataOut]) ZeroValueErrorTransformPresenter(err error) {
	var zeroVal TDataIn
	this.Present(zeroVal, err)
}
