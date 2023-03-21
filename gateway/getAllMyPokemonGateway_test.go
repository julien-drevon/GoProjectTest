package gateway

import (
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2GetPokemonMemory(t *testing.T) {
	assert := assert.New(t)

	repo := Repo{Context: map[string][]domain.Pokemon{"sacha": {{Name: "pikatchu"}, {Name: "tortank"}}}}
	service := GetAllMyPokemonGateway{Context: &repo}

	expected := domain.PokemonsPlayer{Pokemons: []domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, Player: "sacha"}
	actual, _ := service.Get(domain.GetPokemonQuery{Player: "sacha"})

	assert.Equal(expected, actual)
}
