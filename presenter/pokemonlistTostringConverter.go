package presenter

import (
	"clean/core"
	"linq"
	"pokedex/domain"
)

type PokemonListToStringConverter struct {
}

func (converter PokemonListToStringConverter) Convert(data core.PaginationResult[domain.Pokemon]) ([]string, error) {
	return linq.Select(data.Result, func(x domain.Pokemon) string { return x.Name }), nil
}

func NewPokemonlistTostringConverterPresenter() core.TransformPresenter[core.PaginationResult[domain.Pokemon], []string] {
	return core.TransformPresenter[core.PaginationResult[domain.Pokemon], []string]{
		Converter: PokemonListToStringConverter{},
	}
}
