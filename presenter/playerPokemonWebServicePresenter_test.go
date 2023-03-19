package presenter

import (
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemonPaginationWebServicePresenter(t *testing.T) {
	assert := assert.New(t)
	converter := PlayerPokemonWebServicePresenter{}

	expected := domain.PokemonsPlayer{Player: "sacha", Pokemons: []domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}}
	actual, _ := converter.Convert(domain.PokemonsPlayer{Player: "sacha", Pokemons: []domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}})

	assert.Equal(expected, actual)
}
