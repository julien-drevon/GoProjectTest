package gateway

import (
	"linq"
	"pokedex/domain"
)

type AddPokemonGateway struct {
	Context *Repo
}

func (this AddPokemonGateway) Add(query domain.AddPokemonsQuery) (domain.PokemonsPlayer, error) {

	pokeSelect := TransformQueryToPokemon(query)
	AddInContext(pokeSelect, this, query)

	return domain.PokemonsPlayer{Player: query.Player, Pokemons: pokeSelect}, nil
}
func TransformQueryToPokemon(query domain.AddPokemonsQuery) []domain.Pokemon {
	return linq.Select(query.Names, func(x string) domain.Pokemon { return domain.Pokemon{Name: x} })
}

func AddInContext(pokeSelect []domain.Pokemon, this AddPokemonGateway, query domain.AddPokemonsQuery) {
	for _, v := range pokeSelect {
		this.Context.Add(query.Player, v)
	}
}
