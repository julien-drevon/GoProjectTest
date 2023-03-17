package gateway

import (
	"clean/core"
	"pokedex/domain"
)

type GetPokemonReferentielGateway struct {
	PokeList []domain.Pokemon
}

var POKEDEX []domain.Pokemon = []domain.Pokemon{
	{Name: "draco feu"},
	{Name: "pikatchu"},
	{Name: "tortank"},
}

func (this GetPokemonReferentielGateway) Get(query domain.GetPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	return core.NewPaginationResult(POKEDEX, len(POKEDEX), 1, 0), nil
}