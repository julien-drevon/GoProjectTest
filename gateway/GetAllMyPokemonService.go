package gateway

import (
	"clean/core"
	"pokedex/domain"
)

type GetAllMyPokemonService struct {
	PokeList []domain.Pokemon
}

func (service GetAllMyPokemonService) Get(query domain.GetPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	return core.NewPaginationResult(service.PokeList, 2, 1, 0), nil
}
