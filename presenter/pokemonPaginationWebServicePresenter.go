package presenter

import (
	"clean/core"
	"net/http"
	"pokedex/domain"
)

type PokemonPaginationHttpResult struct {
	Data   core.PaginationResult[domain.Pokemon]
	Status int
	Error  string
}

type PokemonPaginationWebServicePresenter struct {
}

func (this PokemonPaginationWebServicePresenter) Convert(data core.PaginationResult[domain.Pokemon], err error) (PokemonPaginationHttpResult, error) {
	if err != nil {
		return PokemonPaginationHttpResult{Data: core.PaginationResult[domain.Pokemon]{}, Status: http.StatusBadRequest, Error: err.Error()}, nil
	}

	return PokemonPaginationHttpResult{Data: data, Status: http.StatusOK}, err
}

// domain.PokemonsPlayer
func NewPokemonPaginationWebServicePresenter() core.TransformPresenter[core.PaginationResult[domain.Pokemon], PokemonPaginationHttpResult] {
	return core.TransformPresenter[core.PaginationResult[domain.Pokemon], PokemonPaginationHttpResult]{
		Converter: PokemonPaginationWebServicePresenter{}.Convert,
	}
}
