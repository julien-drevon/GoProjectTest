package domain

import "clean/core"

type AddPokemonInPokedex struct {
	IAddPokemon
}

func (useCase *AddPokemonInPokedex) Execute(query AddPokemonsQuery, presenter core.IPresentIn[PokemonsPlayer]) {
	pokemons, err := Add(useCase.IAddPokemon, query)
	presenter.Present(pokemons, err)
}

func Add(addProvider IAddPokemon, query AddPokemonsQuery) (PokemonsPlayer, error) {
	return addProvider.Add(query)
}
