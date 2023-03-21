package domain

import "clean/core"

type IAddPokemon interface {
	Add(AddPokemonsQuery) (PokemonsPlayer, error)
}
type IGetPokedemon interface {
	Get(GetPokemonQuery) (PokemonsPlayer, error)
}
type IGetPokedex interface {
	GetPokedex(GetPokemonQuery) (core.PaginationResult[Pokemon], error)
	IsExist(AddPokemonsQuery) bool
}
