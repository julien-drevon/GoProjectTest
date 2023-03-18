package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetAllPokemonsReferentielStub struct {
}

func (service GetAllPokemonsReferentielStub) GetPokedex(query GetPokemonQuery) (PokemonsPlayer, error) {
	return PokemonsPlayer{Player: "sacha", Pokemons: []Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}}, nil
}

func Test_GetAllPokemonsReferentiel(t *testing.T) {
	assert := assert.New(t)

	query := GetPokemonQuery{}
	useCase := GetPokemonReferentiel{IGetPokedex: GetAllPokemonsReferentielStub{}}
	presenter := &core.SimplePresenter[PokemonsPlayer]{}
	useCase.Execute(query, presenter)

	actual, _ := presenter.Print()
	expected := core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 2, 1, 0)

	assert.Equal(expected, actual)
}
