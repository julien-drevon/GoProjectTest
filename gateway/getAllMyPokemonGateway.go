package gateway

import (
	"pokedex/domain"
)

type GetAllMyPokemonGateway struct {
	Context *Repo
}

func (this GetAllMyPokemonGateway) Get(query domain.GetPokemonQuery) (domain.PokemonsPlayer, error) {
	pokeList := this.Context.Get(query.Player, func(p domain.Pokemon) bool { return true })
	return pokeList, nil
}
