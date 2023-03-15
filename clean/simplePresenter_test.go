package core

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplePresenter_Should_PresentData_And_PrintResult(t *testing.T) {
	assert := assert.New(t)
	initialPresenter := SimplePresenter[string]{}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	var compare string = "t"
	presenterIn.Present(compare, nil)
	result, _ := presenterOut.Print()
	assert.Equal(result, compare)
}

func TestSimplePresenter_Should_Return_Error(t *testing.T) {
	assert := assert.New(t)
	initialPresenter := SimplePresenter[string]{}
	var presenterIn IPresentIn[string] = &initialPresenter
	var presenterOut IPresentOut[string] = &initialPresenter
	presenterIn.Present("t", errors.New("ERROR !!!"))
	_, err := presenterOut.Print()
	assert.Equal(err, errors.New("ERROR !!!"))
}
