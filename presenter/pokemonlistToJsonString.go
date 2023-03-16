package presenter

import (
	"clean/core"
	"encoding/json"
	"pokedex/domain"
)

type PokemonListToJsonStringConverter struct {
}

func (converter PokemonListToJsonStringConverter) Convert(data core.PaginationResult[domain.Pokemon]) (string, error) {
	pokeList := data.Result
	emp3, _ := json.Marshal(pokeList)
	return string(emp3), nil
}

func NewPokemonlistToJsonStringPresenter() core.TransformPresenter[core.PaginationResult[domain.Pokemon], string] {
	return core.TransformPresenter[core.PaginationResult[domain.Pokemon], string]{
		Converter: PokemonListToJsonStringConverter{},
	}
}
