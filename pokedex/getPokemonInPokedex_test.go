package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetAllMyPokemonStub struct {
}

func (service GetAllMyPokemonStub) Get(query GetPokemonQuery) (PokemonsPlayer, error) {
	return PokemonsPlayer{Player: "sacha", Pokemons: []Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}}, nil
}

func Test_GetAllMyPokemon(t *testing.T) {
	assert := assert.New(t)

	query := GetPokemonQuery{}
	useCase := GetPokemonInPokedex{IGetPokedemon: GetAllMyPokemonStub{}}
	presenter := &core.SimplePresenter[PokemonsPlayer]{}
	useCase.Execute(query, presenter)

	expected := PokemonsPlayer{Player: "sacha", Pokemons: []Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}}
	actual, _ := presenter.Print()

	assert.Equal(expected, actual)
}
