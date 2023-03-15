package domain

import "clean/core"

type AddPokemonInPokedex struct {
	IAddPokemon
}

func (useCase AddPokemonInPokedex) Execute(query AddPokemonQuery, presenter core.IPresentIn[[]Pokemon]) {
	pokemons, err := Add(useCase.IAddPokemon, query)
	presenter.Present(pokemons, err)
}

func Add(addProvider IAddPokemon, query AddPokemonQuery) ([]Pokemon, error) {
	return addProvider.Add(query)
}
