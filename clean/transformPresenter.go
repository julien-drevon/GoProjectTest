package core

type TransformPresenter[TDataIn any, TDataOut any] struct {
	Result    TDataIn
	Err       error
	Converter IConverting[TDataIn, TDataOut]
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

	return presenter.Converter.Convert(presenter.Result)
}
