package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReferentiel(t *testing.T) {
	//Given
	assert := assert.New(t)
	controller := NewPokemonReferentielControllerJSonAndMemory()
	//when
	presenter := controller.GetReferentiel()
	//Then
	expected := "[{\"name\":\"draco feu\"},{\"name\":\"pikatchu\"},{\"name\":\"tortank\"}]"
	actual, _ := presenter.Print()
	assert.Equal(expected, actual)
}
