package pokemon

import (
	"clean/core"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetPokelist() core.PaginationResult[domain.Pokemon] {
	return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 2, 1, 0)
}

func TestDebile(t *testing.T) {
	assert := assert.New(t)
	converter := PokemonlistTostringConverter{}

	expected := "pikatchu\ntortank"
	actual, _ := converter.Convert(GetPokelist())

	assert.Equal(expected, actual)
}
