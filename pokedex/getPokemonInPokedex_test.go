package domain

import (
	"clean/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetPokemonServiceStub struct {
}

func (service GetPokemonServiceStub) Get(query GetPokemonQuery) ([]Pokemon, error) {
	return []Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, nil
}
func Test_GetAllPokemon(t *testing.T) {
	assert := assert.New(t)

	query := GetPokemonQuery{}
	useCase := GetPokemonInPokedex{IGetPokemon: GetPokemonServiceStub{}}
	presenter := &core.SimplePresenter[[]Pokemon]{}
	wait := []Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}
	useCase.Execute(query, presenter)
	result, _ := presenter.Print()

	assert.Equal(wait, result)
}
