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

	query := GetPokemonQuery{Player: "sacha"}
	useCase := GetPokemonReferentiel{IGetPokedex: GetAllPokemonsReferentielStub{}}
	presenter := &core.SimplePresenter[PokemonsPlayer]{}
	useCase.Execute(query, presenter)

	actual, _ := presenter.Print()
	expected := PokemonsPlayer{Player: "sacha", Pokemons: []Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}}

	assert.Equal(expected, actual)
}
