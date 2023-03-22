package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StringTestConverter struct {
}

func (converter StringTestConverter) Convert(data string, err error) (string, error) {
	if err != nil {
		return "", err

	}
	return data + "1", nil
}

func TestTransformPresenter_Should_PresentStringData_And_PrintStringConvertResult(t *testing.T) {
	assert := assert.New(t)
	converter := StringTestConverter{}
	initialPresenter := TransformPresenter[string, string]{Converter: converter.Convert}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	presenterIn.Present("1", nil)
	actual, _ := presenterOut.Print()
	assert.Equal("11", actual)
}

func TestTransformPresenter_Should_Return_Error_If_UsedCaseError(t *testing.T) {
	assert := assert.New(t)
	converter := StringTestConverter{}
	initialPresenter := TransformPresenter[string, string]{Converter: converter.Convert}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	presenterIn.Present("1", errors.New("Error !!"))
	_, err := presenterOut.Print()
	assert.Equal(err, errors.New("Error !!"))
}

func TestTransformPresenter_Should_Return_Error_IF_ConverterNotInject(t *testing.T) {
	assert := assert.New(t)

	initialPresenter := TransformPresenter[string, string]{Converter: nil}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	presenterIn.Present("1", nil)
	_, err := presenterOut.Print()
	assert.Equal(err, errors.New("Converter must be inject in TransformerPresenter"))
}

func TestTransformPresenter__Return_ZeroValue(t *testing.T) {
	assert := assert.New(t)

	var zeroV string
	initialPresenter := TransformPresenter[string, string]{Converter: nil}

	initialPresenter.ZeroValueErrorTransformPresenter(errors.New("une erreur"))
	actual, errActual := initialPresenter.Print()
	assert.Equal(errors.New("une erreur"), errActual)
	assert.Equal(zeroV, actual)
}
