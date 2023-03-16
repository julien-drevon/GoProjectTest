package domain

import "clean/core"

type GetPokemonQuery struct {
}
type GetPokemonInPokedex struct {
	IGetPokemon
}

func (useCase GetPokemonInPokedex) Execute(query GetPokemonQuery, presenter core.IPresentIn[[]Pokemon]) {
	pokemons, err := Get(useCase.IGetPokemon, query)
	presenter.Present(pokemons, err)
}

func Get(iGetPokemon IGetPokemon, query GetPokemonQuery) ([]Pokemon, error) {
	return iGetPokemon.Get(query)
}
