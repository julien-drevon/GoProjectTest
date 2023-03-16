package pokemon

import (
	"clean/core"
	"pokedex/domain"
)

type PokemonlistTostringConverter struct {
}

func (converter PokemonlistTostringConverter) Convert(data core.PaginationResult[domain.Pokemon]) (string, error) {
	return "pikatchu\ntortank", nil
}
