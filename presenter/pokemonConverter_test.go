package pokemon

import (
	"clean/core"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetPokelist() core.PaginationResult[domain.Pokemon] {
	return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
}

func TestPokemonToString(t *testing.T) {
	assert := assert.New(t)
	converter := PokemonlistTostringConverter{}

	expected := []string{"pikatchu", "tortank"}
	actual, _ := converter.Convert(GetPokelist())

	assert.Equal(expected, actual)
}

func TestPokemonToJson(t *testing.T) {
	assert := assert.New(t)
	converter := PokemonlistToJsonString{}

	expected := "[{\"name\":\"pikatchu\"},{\"name\":\"tortank\"}]"
	actual, _ := converter.Convert(GetPokelist())

	assert.Equal(expected, actual)
}
