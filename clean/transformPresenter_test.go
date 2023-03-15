package core

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type StringTestConverter struct {
}

func getFactory() func(data string) string {
	return func(data string) string { return data + "1" }
}

func TestTransformPresenter_Should_PresentStringData_And_PrintStringConvertResult(t *testing.T) {
	assert := assert.New(t)
	initialPresenter := TransformPresenter[string, string]{factory: getFactory()}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	presenterIn.Present("1", nil)
	val, _ := presenterOut.Print()
	assert.Equal(val, "11")
}

func TestTransformPresenter_Should_Return_Error(t *testing.T) {
	assert := assert.New(t)
	initialPresenter := TransformPresenter[string, string]{factory: getFactory()}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	presenterIn.Present("1", errors.New("Error !!"))
	_, err := presenterOut.Print()
	assert.Equal(err, errors.New("Error !!"))
}
