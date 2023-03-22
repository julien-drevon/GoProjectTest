package presenter

import (
	"errors"
	"net/http"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemonPaginationWebServicePresenter(t *testing.T) {
	assert := assert.New(t)
	converter := PlayerPokemonWebServicePresenter{}

	expected := HttpResponse[domain.PokemonsPlayer]{Status: 200, Error: "", Data: domain.PokemonsPlayer{Player: "sacha", Pokemons: []domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}}}
	actual, _ := converter.Convert(domain.PokemonsPlayer{Player: "sacha", Pokemons: []domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}}, nil)

	assert.Equal(expected, actual)
}

func TestPokemonPaginationWebServicePresenter_error(t *testing.T) {
	assert := assert.New(t)
	converter := PlayerPokemonWebServicePresenter{}

	expected := HttpResponse[domain.PokemonsPlayer]{Status: http.StatusBadRequest, Error: "monuentale erreur"}
	actual, _ := converter.Convert(domain.PokemonsPlayer{Player: "sacha", Pokemons: []domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}}, errors.New("monuentale erreur"))

	assert.Equal(expected, actual)
}
