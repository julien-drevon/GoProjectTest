package domain

import "clean/core"

type GetPokemonQuery struct {
}
type GetPokemonInPokedex struct {
	IGetPokemon
}

func (useCase GetPokemonInPokedex) Execute(query GetPokemonQuery, presenter core.IPresentIn[core.PaginationResult[Pokemon]]) {
	pokemons, err := Get(useCase.IGetPokemon, query)
	presenter.Present(pokemons, err)
}

func Get(iGetPokemon IGetPokemon, query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return iGetPokemon.Get(query)
}
