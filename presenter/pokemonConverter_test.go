package presenter

import (
	"clean/core"
	"errors"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetPokelist() core.PaginationResult[domain.Pokemon] {
	return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
}

func TestPokemonListToString(t *testing.T) {
	assert := assert.New(t)
	converter := PokemonListToStringConverter{}

	expected := []string{"pikatchu", "tortank"}
	actual, _ := converter.Convert(GetPokelist(), nil)

	assert.Equal(expected, actual)

	errExpected := errors.New("")
	_, errActual := converter.Convert(GetPokelist(), errors.New(""))
	assert.Equal(errExpected, errActual)
}

func TestPokemonListToJson(t *testing.T) {
	assert := assert.New(t)
	converter := PokemonListToJsonStringConverter{}

	expected := "[{\"name\":\"pikatchu\"},{\"name\":\"tortank\"}]"
	actual, _ := converter.Convert(GetPokelist(), nil)

	assert.Equal(expected, actual)

	errExpected := errors.New("")
	_, errActual := converter.Convert(GetPokelist(), errors.New(""))
	assert.Equal(errExpected, errActual)
}

func TestPokemonToJson(t *testing.T) {
	assert := assert.New(t)
	converter := PokemonListToJsonStringConverter{}

	expected := "[{\"name\":\"pikatchu\"},{\"name\":\"tortank\"}]"
	actual, _ := converter.Convert(GetPokelist(), nil)

	assert.Equal(expected, actual)

	errExpected := errors.New("")
	_, errActual := converter.Convert(GetPokelist(), errors.New(""))
	assert.Equal(errExpected, errActual)
}
