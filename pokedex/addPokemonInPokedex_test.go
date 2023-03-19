package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AddOnePokemonServiceStub struct {
}

func (service AddOnePokemonServiceStub) Add(query AddPokemonsQuery) (PokemonsPlayer, error) {
	return PokemonsPlayer{Player: "sacha", Pokemons: []Pokemon{{Name: "pikatch"}}}, nil
}

func Test_AddOnePokemon(t *testing.T) {
	assert := assert.New(t)

	query := AddPokemonsQuery{Player: "sacha", Names: []string{"pikatchu"}}
	useCase := AddPokemonInPokedex{IAddPokemon: AddOnePokemonServiceStub{}}
	presenter := &core.SimplePresenter[PokemonsPlayer]{}
	useCase.Execute(query, presenter)

	actual, _ := presenter.Print()
	expected := PokemonsPlayer{Player: "sacha", Pokemons: []Pokemon{{Name: "pikatch"}}}

	assert.Equal(expected, actual)
}
