package domain

import "clean/core"

type GetPokemonReferentiel struct {
	IGetPokedex
}

func (useCase GetPokemonReferentiel) Execute(query GetPokemonQuery, presenter core.IPresentIn[PokemonsPlayer]) {
	pokemons, err := useCase.Get(useCase.IGetPokedex, query)
	presenter.Present(pokemons, err)
}

func (useCase GetPokemonReferentiel) Get(iGetPokemon IGetPokedex, query GetPokemonQuery) (PokemonsPlayer, error) {
	return iGetPokemon.GetPokedex(query)
}
