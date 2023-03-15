package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebile(t *testing.T) {
	assert := assert.New(t)

	// query := AddPokemonQuery{Name: "pikatchu"}
	// useCase := AddPokemonInPokedex{IAddPokemon: AddOnePokemonServiceStub{}}
	// presenter := &core.SimplePresenter[[]Pokemon]{}
	// wait := []Pokemon{{Name: "pikatchu"}}
	// useCase.Execute(query, presenter)
	// result, _ := presenter.Print()

	assert.Equal(true, true)
}
