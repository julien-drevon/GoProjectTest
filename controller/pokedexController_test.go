package controller

import (
	//"clean/core"
	// "pokedex/domain"
	// "presenter"

	"gateway"
	"presenter"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func GetDebile() core.PaginationResult[domain.Pokemon] {
// 	return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
// }

func TestGetPokemonIntegration(t *testing.T) {
	assert := assert.New(t)

	controller := PokedexController[[]string]{Presenter: presenter.NewPokemonlistTostringConverterPresenter(), Gateway: gateway.GetAllMyPokemonService{}}
	presenter := controller.GetMyPokemons()

	expected := []string{"pikatchu", "tortank"}
	actual, _ := presenter.Print()

	assert.Equal(expected, actual)
}
