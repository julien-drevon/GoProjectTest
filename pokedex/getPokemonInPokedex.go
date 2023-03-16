package domain

import "clean/core"

type GetPokemonQuery struct {
}

type GetPokemonInPokedex struct {
	IGetPokedemon
}

func (useCase GetPokemonInPokedex) Execute(query GetPokemonQuery, presenter core.IPresentIn[core.PaginationResult[Pokemon]]) {
	pokemons, err := useCase.Get(useCase.IGetPokedemon, query)
	presenter.Present(pokemons, err)
}

func (useCase GetPokemonInPokedex) Get(iGetPokemon IGetPokedemon, query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return iGetPokemon.Get(query)
}
