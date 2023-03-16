package domain

import "clean/core"

type IAddPokemon interface {
	Add(AddPokemonQuery) ([]Pokemon, error)
}
type IGetPokemon interface {
	Get(GetPokemonQuery) (core.PaginationResult[Pokemon], error)
}
