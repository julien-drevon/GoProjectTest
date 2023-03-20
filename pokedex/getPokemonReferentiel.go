package domain

import "clean/core"

type GetPokemonReferentiel struct {
	IGetPokedex
}

func (this GetPokemonReferentiel) Execute(query GetPokemonQuery, presenter core.IPresentIn[core.PaginationResult[Pokemon]]) {
	pokemons, err := this.Get(this.IGetPokedex, query)
	presenter.Present(pokemons, err)
}

func (this GetPokemonReferentiel) Get(iGetPokemon IGetPokedex, query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return iGetPokemon.GetPokedex(query)
}
