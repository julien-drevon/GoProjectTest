package gateway

import (
	"clean/core"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2GetPokemonMemory(t *testing.T) {
	assert := assert.New(t)

	service := GetAllMyPokemonGateway{PokeList: []domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}}
	expected := core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
	actual, _ := service.Get(domain.GetPokemonQuery{})

	assert.Equal(expected, actual)
}
