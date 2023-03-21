package gateway

import (
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemonReferentialFile(t *testing.T) {
	assert := assert.New(t)

	service := GetPokemonReferentialFleGateway{}
	expected := 4
	actual, err := service.GetPokedex(domain.GetPokemonQuery{})
	assert.Equal(nil, err)
	assert.Equal(expected, actual.Pagination.Total)
}

func Test_IsAny_GetPokemonReferentialFile(t *testing.T) {
	assert := assert.New(t)

	service := GetPokemonReferentialFleGateway{}

	actual := service.IsExist(domain.AddPokemonsQuery{Names: []string{"pikatchu", "bulbizare"}})
	assert.Equal(true, actual)
	actual = service.IsExist(domain.AddPokemonsQuery{Names: []string{"toto"}})
	assert.Equal(false, actual)
}
