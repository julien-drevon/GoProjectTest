package controller

import (
	//"clean/core"
	// "pokedex/domain"
	// "presenter"

	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func GetDebile() core.PaginationResult[domain.Pokemon] {
// 	return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
// }

func TestGetPokemonIntegration(t *testing.T) {
	assert := assert.New(t)
	controller := NewControllerJSonAndMemory(make([]domain.Pokemon, 0))
	presenter := controller.GetMyPokemons()

	expected := "[]"
	actual, _ := presenter.Print()

	assert.Equal(expected, actual)
}

// func TestAddPokemonIntegration(t *testing.T) {
// 	assert := assert.New(t)
// 	controller := NewControllerJSonAndMemory(make([]domain.Pokemon, 0))
// 	presenter := controller.AddPokemon("pikatch")

// 	expected := "[{\"name\":\"pikatchu\"},{\"name\":\"tortank\"}]"
// 	actual, _ := presenter.Print()

// 	assert.Equal(expected, actual)
// }
