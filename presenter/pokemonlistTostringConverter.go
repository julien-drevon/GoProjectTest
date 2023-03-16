package pokemon

import (
	"clean/core"
	"linq"
	"pokedex/domain"
)

type PokemonlistTostringConverter struct {
}

func (converter PokemonlistTostringConverter) Convert(data core.PaginationResult[domain.Pokemon]) ([]string, error) {
	return linq.Select(data.Result, func(x domain.Pokemon) string { return x.Name }), nil
}
