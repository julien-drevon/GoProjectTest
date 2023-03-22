package presenter

import (
	"clean/core"
	"linq"
	"pokedex/domain"
)

type PokemonListToStringConverter struct {
}

func (this PokemonListToStringConverter) Convert(data core.PaginationResult[domain.Pokemon], err error) ([]string, error) {
	return linq.Select(data.Result, func(x domain.Pokemon) string { return x.Name }), err
}

func NewPokemonlistTostringConverterPresenter() core.TransformPresenter[core.PaginationResult[domain.Pokemon], []string] {
	return core.TransformPresenter[core.PaginationResult[domain.Pokemon], []string]{
		Converter: PokemonListToStringConverter{}.Convert,
	}
}
