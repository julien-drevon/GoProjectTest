package gateway

import (
	"clean/core"
	"pokedex/domain"
)

type GetAllMyPokemonGateway struct {
	Context *Repo[domain.Pokemon]
}

func (this GetAllMyPokemonGateway) Get(query domain.GetPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	pokeList := this.Context.Get()
	return core.NewPaginationResult(pokeList, len(pokeList), 1, 0), nil
}
