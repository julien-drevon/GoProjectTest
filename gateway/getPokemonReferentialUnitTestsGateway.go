package gateway

import (
	"clean/core"
	"linq"
	"pokedex/domain"
)

type GetPokemonReferentialUnitTestsGateway struct {
}

var POKEDEX []domain.Pokemon = []domain.Pokemon{
	{Name: "draco feu"},
	{Name: "pikatchu"},
	{Name: "tortank"},
}

func (this GetPokemonReferentialUnitTestsGateway) GetPokedex(query domain.GetPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	return core.NewPaginationResult(POKEDEX, len(POKEDEX), 1, 0), nil
}
func (this GetPokemonReferentialUnitTestsGateway) IsExist(query domain.AddPokemonsQuery) bool {
	return linq.Any(POKEDEX, func(x domain.Pokemon) bool { return linq.Any(query.Names, func(s string) bool { return s == x.Name }) })
}
