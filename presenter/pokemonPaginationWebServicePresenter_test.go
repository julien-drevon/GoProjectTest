package presenter

import (
	"clean/core"
	"errors"
	"net/http"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerPokemonWebServicePresenter_ShouldReturn_HttpOk_Struct(t *testing.T) {
	assert := assert.New(t)
	converter := PokemonPaginationWebServicePresenter{}

	expected := PokemonPaginationHttpResult{Data: core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 4, 1, 0), Status: http.StatusOK}
	actual, _ := converter.Convert(core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 4, 1, 0), nil)

	assert.Equal(expected, actual)
}
func TestPlayerPokemonWebServicePresenter_Should_Return_HttpError_Struc(t *testing.T) {
	assert := assert.New(t)
	converter := PokemonPaginationWebServicePresenter{}

	expected := PokemonPaginationHttpResult{Data: core.PaginationResult[domain.Pokemon]{}, Error: "Une erreure s'est produite", Status: http.StatusBadRequest}
	actual, _ := converter.Convert(core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 4, 1, 0), errors.New("Une erreure s'est produite"))

	assert.Equal(expected, actual)
}
