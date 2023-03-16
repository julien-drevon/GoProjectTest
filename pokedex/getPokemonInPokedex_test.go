package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetPokemonServiceStub struct {
}

func (service GetPokemonServiceStub) Get(query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 2, 1, 0), nil
}
func Test_GetAllMyPokemon(t *testing.T) {
	assert := assert.New(t)

	query := GetPokemonQuery{}
	useCase := GetPokemonInPokedex{IGetPokemon: GetPokemonServiceStub{}}
	presenter := &core.SimplePresenter[core.PaginationResult[Pokemon]]{}
	wait := core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 2, 1, 0)
	useCase.Execute(query, presenter)
	result, _ := presenter.Print()

	assert.Equal(wait, result)
}
