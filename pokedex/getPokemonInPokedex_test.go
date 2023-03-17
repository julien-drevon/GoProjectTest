package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetAllMyPokemonStub struct {
}

func (service GetAllMyPokemonStub) Get(query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 2, 1, 0), nil
}

func Test_GetAllMyPokemon(t *testing.T) {
	assert := assert.New(t)

	query := GetPokemonQuery{}
	useCase := GetPokemonInPokedex{IGetPokedemon: GetAllMyPokemonStub{}}
	presenter := &core.SimplePresenter[core.PaginationResult[Pokemon]]{}
	useCase.Execute(query, presenter)

	expected := core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 2, 1, 0)
	actual, _ := presenter.Print()

	assert.Equal(expected, actual)
}
