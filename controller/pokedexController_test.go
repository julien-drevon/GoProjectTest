package controller

import (
	"errors"
	"gateway"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemon(t *testing.T) {
	assert := assert.New(t)
	controller := NewControllerTestWithInit(map[string][]domain.Pokemon{"sacha": {{Name: "pikatchu"}}})
	presenter := controller.GetMyPokemons("sacha")

	expected := "{\"Pokemons\":[{\"name\":\"pikatchu\"}],\"Player\":\"sacha\"}"
	actual, _ := presenter.Print()

	assert.Equal(expected, actual)
}

func TestAddWithEmptyName(t *testing.T) {
	assert := assert.New(t)

	controller := NewControllerTest()
	presenter := controller.AddPokemons("", []string{"pikatchu"})

	//expected := "[{\"name\":\"pikatchu\"}]"
	expected := errors.New("player should not be empty")
	_, actual := presenter.Print()

	assert.Equal(expected, actual)
}

func TestAddAndGetPokemon(t *testing.T) {
	//Given
	assert := assert.New(t)
	controller := NewControllerTest()
	//when
	controller.AddPokemons("sacha", []string{"pikatchu"})
	presenter := controller.GetMyPokemons("sacha")
	//Then
	expected := "{\"Pokemons\":[{\"name\":\"pikatchu\"}],\"Player\":\"sacha\"}"
	actual, _ := presenter.Print()
	assert.Equal(expected, actual)
}

func TestAddPokemonIntegrationWithEmptyName(t *testing.T) {
	//Given
	assert := assert.New(t)
	controller := NewControllerTest()
	//when
	presenter := controller.AddPokemons("", []string{"pikatchu"})
	//Then
	expected := errors.New("player should not be empty")
	_, actual := presenter.Print()
	assert.Equal(expected, actual)
}

func TestAddAddPokemonNotInReferential(t *testing.T) {
	//Given
	assert := assert.New(t)
	controller := NewControllerTest()
	//when
	presenter := controller.AddPokemons("sacha", []string{"toto"})

	expected, errExpected := "", errors.New("Pokemon Not In Referential could not be add")
	actual, err := presenter.Print()

	assert.Equal(expected, actual)
	assert.Equal(errExpected, err)
}

func NewControllerTest() PokedexController[string] {
	repo := gateway.NewRepo()
	return NewControllerJSonAndMemory(repo)
}

func NewControllerTestWithInit(buf map[string][]domain.Pokemon) PokedexController[string] {
	repo := gateway.Repo{Context: buf}
	return NewControllerJSonAndMemory(repo)
}
