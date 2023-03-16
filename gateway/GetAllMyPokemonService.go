package gateway

import (
	"clean/core"
	"pokedex/domain"
)

type GetAllMyPokemonService struct {
}

func (service GetAllMyPokemonService) Get(query domain.GetPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0), nil
}
