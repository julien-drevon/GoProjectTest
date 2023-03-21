package domain

import "clean/core"

type GetPokemonInPokedex struct {
	IGetPokedemon
}

func (this GetPokemonInPokedex) Execute(query GetPokemonQuery, presenter core.IPresentIn[PokemonsPlayer]) {
	pokemons, err := this.Get(this.IGetPokedemon, query)
	presenter.Present(pokemons, err)
}

func (this GetPokemonInPokedex) Get(iGetPokemon IGetPokedemon, query GetPokemonQuery) (PokemonsPlayer, error) {
	return iGetPokemon.Get(query)
}
