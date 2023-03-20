package domain

import "clean/core"

type GetPokemonReferential struct {
	IGetPokedex
}

func (this GetPokemonReferential) Execute(query GetPokemonQuery, presenter core.IPresentIn[core.PaginationResult[Pokemon]]) {
	pokemons, err := this.Get(this.IGetPokedex, query)
	presenter.Present(pokemons, err)
}

func (this GetPokemonReferential) Get(iGetPokemon IGetPokedex, query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
	return iGetPokemon.GetPokedex(query)
}
