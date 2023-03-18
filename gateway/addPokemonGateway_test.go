package gateway

import (
	"errors"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPokemonWithEmptyName(t *testing.T) {
	assert := assert.New(t)
	repo := NewRepo()
	service := AddPokemonGateway{Context: &repo}
	expected := errors.New("Player should not be empty")
	_, actual := service.Add(domain.AddPokemonsQuery{Names: []string{"pikatchu", "tortank"}})

	assert.Equal(expected, actual)
}
func Test2AddPokemonMemory(t *testing.T) {
	assert := assert.New(t)
	repo := NewRepo()
	service := AddPokemonGateway{Context: &repo}
	expected := domain.PokemonsPlayer{Player: "sacha", Pokemons: []domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}}
	actual, _ := service.Add(domain.AddPokemonsQuery{Player: "sacha", Names: []string{"pikatchu", "tortank"}})

	assert.Equal(expected, actual)
}
func Test2AddPokemonMemoryWithOther(t *testing.T) {
	assert := assert.New(t)
	repo := NewRepoWithContext(map[string][]domain.Pokemon{"sacha": {{Name: "draco feu"}}})
	service := AddPokemonGateway{Context: &repo}

	service.Add(domain.AddPokemonsQuery{Player: "sacha", Names: []string{"pikatchu", "tortank"}})

	expected := domain.PokemonsPlayer{Player: "sacha", Pokemons: []domain.Pokemon{{Name: "draco feu"}, {Name: "pikatchu"}, {Name: "tortank"}}}
	actual := service.Context.Get("sacha", func(p domain.Pokemon) bool { return true })

	assert.Equal(expected, actual)
}
