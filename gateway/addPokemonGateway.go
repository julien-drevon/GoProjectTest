package gateway

import (
	"errors"
	"linq"
	"pokedex/domain"
	"strings"
)

type AddPokemonGateway struct {
	Context *Repo
}

func (this AddPokemonGateway) Add(query domain.AddPokemonsQuery) (domain.PokemonsPlayer, error) {
	if strings.Trim(query.Player, " ") == "" {
		var defaut domain.PokemonsPlayer
		return defaut, errors.New("Player should not be empty")
	}
	pokeSelect := linq.Select(query.Names, func(x string) domain.Pokemon { return domain.Pokemon{Name: x} })
	for _, v := range pokeSelect {
		this.Context.Add(query.Player, v)
	}

	return domain.PokemonsPlayer{Player: query.Player, Pokemons: pokeSelect}, nil
}
