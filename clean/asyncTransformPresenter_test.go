package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (converter StringTestConverter) ConvertAsync(data string, c chan string, e chan error) {
	c <- data + "1"
	e <- nil
}

func Test_AsyncTransformPresenter_Should_PresentStringData_And_PrintStringConvertResult(t *testing.T) {
	assert := assert.New(t)
	converter := StringTestConverter{}
	initialPresenter := AsyncTransformPresenter[string, string]{Converter: converter}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	presenterIn.Present("1", nil)
	actual, _ := presenterOut.Print()
	assert.Equal("11", actual)
}

func TestAsyncTransformPresenter_Should_Return_Error_If_UsedCaseError(t *testing.T) {
	assert := assert.New(t)
	converter := StringTestConverter{}
	initialPresenter := AsyncTransformPresenter[string, string]{Converter: converter}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	presenterIn.Present("1", errors.New("Error !!"))
	_, err := presenterOut.Print()
	assert.Equal(err, errors.New("Error !!"))
}

func TestAsyncTransformPresenter_Should_Return_Error_IF_ConverterNotInject(t *testing.T) {
	assert := assert.New(t)

	initialPresenter := AsyncTransformPresenter[string, string]{Converter: nil}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	presenterIn.Present("1", nil)
	_, err := presenterOut.Print()
	assert.Equal(err, errors.New("Converter must be inject in TransformerPresenter"))
}
