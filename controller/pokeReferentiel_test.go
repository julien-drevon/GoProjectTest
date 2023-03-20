package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReferential(t *testing.T) {
	//Given
	assert := assert.New(t)
	controller := NewPokemonReferentialControllerJSonAndMemory()
	//when
	presenter := controller.GetReferential()
	//Then
	expected := "[{\"name\":\"draco feu\"},{\"name\":\"pikatchu\"},{\"name\":\"tortank\"}]"
	actual, _ := presenter.Print()
	assert.Equal(expected, actual)
}
