package presenter

import (
	"clean/core"
	"encoding/json"
	"pokedex/domain"
)

type PokemonListToJsonStringConverter struct {
}

type PokemonPlayerToJsonStringConverter struct {
}

func (this PokemonListToJsonStringConverter) Convert(data core.PaginationResult[domain.Pokemon], err error) (string, error) {
	if err != nil {
		return "", err
	}
	pokeList := data.Result
	emp3, err := json.Marshal(pokeList)
	if err != nil {
		return "", err
	}
	return string(emp3), nil
}

func (this PokemonPlayerToJsonStringConverter) Convert(data domain.PokemonsPlayer, err error) (string, error) {
	if err != nil {
		return "", err
	}
	pokeList := data
	emp3, err := json.Marshal(pokeList)
	if err != nil {
		return "", err
	}
	return string(emp3), nil
}

func NewPokemonListToJsonStringPresenter() core.TransformPresenter[core.PaginationResult[domain.Pokemon], string] {
	return core.TransformPresenter[core.PaginationResult[domain.Pokemon], string]{
		Converter: PokemonListToJsonStringConverter{}.Convert}
}

func NewPokemonPlayerToJsonStringPresenter() core.TransformPresenter[domain.PokemonsPlayer, string] {
	return core.TransformPresenter[domain.PokemonsPlayer, string]{
		Converter: PokemonPlayerToJsonStringConverter{}.Convert}
}

func NewPaginedListPokemonToJsonStringPresenter() core.TransformPresenter[domain.PokemonsPlayer, string] {
	return core.TransformPresenter[domain.PokemonsPlayer, string]{
		Converter: PokemonPlayerToJsonStringConverter{}.Convert}
}
