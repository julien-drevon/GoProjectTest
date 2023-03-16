package domain

import "clean/core"

type GetPokemonDefinition struct {
	IGetPokedex
}

func (useCase GetPokemonDefinition) Execute(query GetPokemonQuery, presenter core.IPresentIn[core.PaginationResult[Pokemon]]) {
	pokemons, err := useCase.Get(useCase.IGetPokedex, query)
	presenter.Present(pokemons, err)
}

func (useCase GetPokemonDefinition) Get(iGetPokemon IGetPokedex, query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return iGetPokemon.Get(query)
}
