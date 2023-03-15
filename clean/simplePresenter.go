package core

type SimplePresenter[TData any] struct {
	Result TData
	Err    error
}

func (presenter *SimplePresenter[TData]) Present(data TData, err error) {
	presenter.Result = data
	presenter.Err = err
}

func (presenter SimplePresenter[TData]) Print() (TData, error) {
	return presenter.Result, presenter.Err
}
