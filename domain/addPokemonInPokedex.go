package domain

import (
	"clean/core"
	"errors"
)

const POKEMON_NOT_EXIST_NOT_Be_ADDED = "Pokemon Not In Referential could not be add"

type AddPokemonInPokedex struct {
	IAddPokemon
	IGetPokedex
}

func (this *AddPokemonInPokedex) Execute(query AddPokemonsQuery, presenter core.IPresentIn[PokemonsPlayer]) {
	var pokemons PokemonsPlayer
	isPokemonExist, err := this.IsExistPokemon(query)
	if isPokemonExist {
		pokemons, err = this.Add(query)
	}
	presenter.Present(pokemons, err)
}

func (this AddPokemonInPokedex) IsExistPokemon(query AddPokemonsQuery) (bool, error) {
	if this.IGetPokedex.IsExist(query) {
		return true, nil
	}

	return false, errors.New(POKEMON_NOT_EXIST_NOT_Be_ADDED)
}

func (this *AddPokemonInPokedex) Add(query AddPokemonsQuery) (PokemonsPlayer, error) {
	return this.IAddPokemon.Add(query)
}
