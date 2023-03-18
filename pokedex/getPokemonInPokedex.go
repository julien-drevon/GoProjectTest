package domain

import "clean/core"

type GetPokemonInPokedex struct {
	IGetPokedemon
}

func (useCase GetPokemonInPokedex) Execute(query GetPokemonQuery, presenter core.IPresentIn[PokemonsPlayer]) {
	pokemons, err := useCase.Get(useCase.IGetPokedemon, query)
	presenter.Present(pokemons, err)
}

func (useCase GetPokemonInPokedex) Get(iGetPokemon IGetPokedemon, query GetPokemonQuery) (PokemonsPlayer, error) {
	return iGetPokemon.Get(query)
}
