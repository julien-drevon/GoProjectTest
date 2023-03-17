package gateway

import (
	"clean/core"
	"pokedex/domain"
)

type GetAllMyPokemonService struct {
	PokeList []domain.Pokemon
}

func (this GetAllMyPokemonService) Get(query domain.GetPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	return core.NewPaginationResult(this.PokeList, len(this.PokeList), 1, 0), nil
}
