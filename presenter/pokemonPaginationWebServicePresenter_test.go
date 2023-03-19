package presenter

import (
	"clean/core"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerPokemonWebServicePresenter(t *testing.T) {
	assert := assert.New(t)
	converter := PokemonPaginationWebServicePresenter{}

	expected := core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 4, 1, 0)
	actual, _ := converter.Convert(core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 4, 1, 0))

	assert.Equal(expected, actual)
}
