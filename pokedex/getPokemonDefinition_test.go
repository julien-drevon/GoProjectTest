package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetAllPokemonsDefinition struct {
}

func (service GetAllPokemonsDefinition) Get(query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 2, 1, 0), nil
}

func Test_GetAllPokemonsDefinition(t *testing.T) {
	assert := assert.New(t)

	query := GetPokemonQuery{}
	useCase := GetPokemonDefinition{IGetPokedex: GetAllPokemonsDefinition{}}
	presenter := &core.SimplePresenter[core.PaginationResult[Pokemon]]{}
	wait := core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 2, 1, 0)
	useCase.Execute(query, presenter)
	result, _ := presenter.Print()

	assert.Equal(wait, result)
}
