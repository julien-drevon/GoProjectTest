package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AddOnePokemonServiceStub struct {
}

func (service AddOnePokemonServiceStub) Add(query AddPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return core.NewPaginationResult([]Pokemon{{Name: "pikatch"}}, 1, 1, 0), nil
}

func Test_AddOnePokemon(t *testing.T) {
	assert := assert.New(t)

	query := AddPokemonQuery{Name: "pikatchu"}
	useCase := AddPokemonInPokedex{IAddPokemon: AddOnePokemonServiceStub{}}
	presenter := &core.SimplePresenter[core.PaginationResult[Pokemon]]{}
	wait := core.NewPaginationResult([]Pokemon{{Name: "pikatch"}}, 1, 1, 0)
	useCase.Execute(query, presenter)
	result, _ := presenter.Print()

	assert.Equal(wait, result)
}
