package presenter

import (
	"clean/core"
	"net/http"
	"pokedex/domain"
)

type PokemonPaginationWebServicePresenter struct {
}

func (this PokemonPaginationWebServicePresenter) Convert(data core.PaginationResult[domain.Pokemon], err error) (PokemonPaginationHttpResult[core.PaginationResult[domain.Pokemon]], error) {
	if err != nil {
		return PokemonPaginationHttpResult[core.PaginationResult[domain.Pokemon]]{Data: core.PaginationResult[domain.Pokemon]{}, Status: http.StatusBadRequest, Error: err.Error()}, nil
	}

	return PokemonPaginationHttpResult[core.PaginationResult[domain.Pokemon]]{Data: data, Status: http.StatusOK}, err
}

// domain.PokemonsPlayer
func NewPokemonPaginationWebServicePresenter() core.TransformPresenter[core.PaginationResult[domain.Pokemon], PokemonPaginationHttpResult[core.PaginationResult[domain.Pokemon]]] {
	return core.TransformPresenter[core.PaginationResult[domain.Pokemon], PokemonPaginationHttpResult[core.PaginationResult[domain.Pokemon]]]{
		Converter: PokemonPaginationWebServicePresenter{}.Convert,
	}
}
