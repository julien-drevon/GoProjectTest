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

	query := AddPokemonQuery{Names: []string{"ikatchu"}}
	useCase := AddPokemonInPokedex{IAddPokemon: AddOnePokemonServiceStub{}}
	presenter := &core.SimplePresenter[core.PaginationResult[Pokemon]]{}
	useCase.Execute(query, presenter)

	actual, _ := presenter.Print()
	expected := core.NewPaginationResult([]Pokemon{{Name: "pikatch"}}, 1, 1, 0)

	assert.Equal(expected, actual)
}
