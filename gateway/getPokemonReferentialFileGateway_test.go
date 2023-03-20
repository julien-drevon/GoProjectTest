package gateway

import (
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemonReferentialFile(t *testing.T) {
	assert := assert.New(t)

	service := GetPokemonReferentialFleGateway{}
	expected := 3
	actual, err := service.GetPokedex(domain.GetPokemonQuery{})
	assert.Equal(nil, err)
	assert.Equal(expected, actual.Pagination.Total)
}