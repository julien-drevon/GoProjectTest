package gateway

import (
	"clean/core"
	"linq"
	"pokedex/domain"
)

type AddPokemonService struct {
	PokeList []domain.Pokemon
}

func (this *AddPokemonService) Add(query domain.AddPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	pokeSelect := linq.Select(query.Names, func(x string) domain.Pokemon { return domain.Pokemon{Name: x} })
	for _, v := range pokeSelect {
		this.PokeList = append(this.PokeList, v)
	}

	return core.NewPaginationResult(pokeSelect, len(pokeSelect), 1, 0), nil
}
