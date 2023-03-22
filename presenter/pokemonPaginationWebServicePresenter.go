package presenter

import (
	"clean/core"
	"net/http"
	"pokedex/domain"
)

type PokemonPaginationHttpResult struct {
	Data   core.PaginationResult[domain.Pokemon]
	STATUS int
}

type PokemonPaginationWebServicePresenter struct {
}

func (this PokemonPaginationWebServicePresenter) Convert(data core.PaginationResult[domain.Pokemon]) (PokemonPaginationHttpResult, error) {
	return PokemonPaginationHttpResult{Data: data, STATUS: http.StatusOK}, nil
}

// domain.PokemonsPlayer
func NewPokemonPaginationWebServicePresenter() core.TransformPresenter[core.PaginationResult[domain.Pokemon], PokemonPaginationHttpResult] {
	return core.TransformPresenter[core.PaginationResult[domain.Pokemon], PokemonPaginationHttpResult]{
		Converter: PokemonPaginationWebServicePresenter{}.Convert,
	}
}
