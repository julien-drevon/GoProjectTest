package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetAllPokemonsReferentielStub struct {
}

func (service GetAllPokemonsReferentielStub) GetPokedex(query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 4, 1, 0), nil
}

func Test_GetAllPokemonsReferentiel(t *testing.T) {
	assert := assert.New(t)

	query := GetPokemonQuery{Player: "sacha"}
	useCase := GetPokemonReferentiel{IGetPokedex: GetAllPokemonsReferentielStub{}}
	presenter := &core.SimplePresenter[core.PaginationResult[Pokemon]]{}
	useCase.Execute(query, presenter)

	actual, _ := presenter.Print()
	expected := core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 4, 1, 0)

	assert.Equal(expected, actual)
}
