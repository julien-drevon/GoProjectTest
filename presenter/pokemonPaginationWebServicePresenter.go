package presenter

import (
	"clean/core"
	"pokedex/domain"
)

type PokemonPaginationWebServicePresenter struct {
}

func (this PokemonPaginationWebServicePresenter) Convert(data core.PaginationResult[domain.Pokemon]) (core.PaginationResult[domain.Pokemon], error) {
	return data, nil
}

// domain.PokemonsPlayer
func NewPokemonPaginationWebServicePresenter() core.TransformPresenter[core.PaginationResult[domain.Pokemon], core.PaginationResult[domain.Pokemon]] {
	return core.TransformPresenter[core.PaginationResult[domain.Pokemon], core.PaginationResult[domain.Pokemon]]{
		Converter: PokemonPaginationWebServicePresenter{}.Convert,
	}
}
