package gateway

import (
	"clean/core"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetdPokemonReferentiel(t *testing.T) {
	assert := assert.New(t)

	service := GetPokemonReferentielGateway{}
	expected := core.NewPaginationResult([]domain.Pokemon{{Name: "draco feu"}, {Name: "pikatchu"}, {Name: "tortank"}}, 3, 1, 0)
	actual, _ := service.Get(domain.GetPokemonQuery{})

	assert.Equal(expected, actual)
}