package controller

import (
	"gateway"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemonIntegration(t *testing.T) {
	assert := assert.New(t)
	controller := NewControllerTestWithInit(map[string][]domain.Pokemon{"sacha": {{Name: "pikatchu"}}})
	presenter := controller.GetMyPokemons()

	expected := "[{\"name\":\"pikatchu\"}]"
	actual, _ := presenter.Print()

	assert.Equal(expected, actual)
}

func TestAddWithEmptyName(t *testing.T) {
	assert := assert.New(t)

	controller := NewControllerTest()
	presenter := controller.AddPokemons([]string{"pikatchu"})

	//expected := "[{\"name\":\"pikatchu\"}]"
	expected := "name should not be empty"
	_, actual := presenter.Print()

	assert.Equal(expected, actual)
}

func TestAddAndGetPokemonIntegration(t *testing.T) {
	//Given
	assert := assert.New(t)
	controller := NewControllerTest()
	//when
	controller.AddPokemons([]string{"pikatchu"})
	presenter := controller.GetMyPokemons()
	//Then
	expected := "[{\"name\":\"pikatchu\"}]"
	actual, _ := presenter.Print()
	assert.Equal(expected, actual)
}

func NewControllerTest() PokedexController[string] {
	repo := gateway.NewRepo()
	return NewControllerJSonAndMemory(repo)
}

func NewControllerTestWithInit(buf map[string][]domain.Pokemon) PokedexController[string] {
	repo := gateway.Repo{Context: buf}
	return NewControllerJSonAndMemory(repo)
}
