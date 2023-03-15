package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AddOnePokemonServiceStub struct {
}

func (service AddOnePokemonServiceStub) Add(query AddPokemonQuery) ([]Pokemon, error) {
	return []Pokemon{{Name: "pikatchu"}}, nil
}

func Test_AddOnePokemon(t *testing.T) {
	assert := assert.New(t)

	query := AddPokemonQuery{Name: "pikatchu"}
	useCase := AddPokemonInPokedex{IAddPokemon: AddOnePokemonServiceStub{}}
	presenter := &core.SimplePresenter[[]Pokemon]{}
	wait := []Pokemon{{Name: "pikatchu"}}
	useCase.Execute(query, presenter)
	result, _ := presenter.Print()

	assert.Equal(wait, result)
}
