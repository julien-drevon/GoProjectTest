package gateway

import (
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2AddPokemonMemory(t *testing.T) {
	assert := assert.New(t)
	repo := NewRepoForUnitTests()
	service := AddPokemonGateway{Context: &repo}
	expected := domain.PokemonsPlayer{Player: "sacha", Pokemons: []domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}}
	actual, _ := service.Add(domain.AddPokemonsQuery{Player: "sacha", Names: []string{"pikatchu", "tortank"}})

	assert.Equal(expected, actual)
}
