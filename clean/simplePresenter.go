package core

type SimplePresenter[TData any] struct {
	Result TData
	Err    error
}

func (this *SimplePresenter[TData]) Present(data TData, err error) {
	this.Result = data
	this.Err = err
}

func (this SimplePresenter[TData]) Print() (TData, error) {
	return this.Result, this.Err
}
