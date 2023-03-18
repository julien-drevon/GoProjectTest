package domain

import (
	"errors"
	"strings"
)

//	type IExecuteUseCase[TQuery any, TData any] interface {
//		Execute(TQuery, IPresentIn[TData])
//	}
type AddPokemonsQuery struct {
	Names  []string `json:"names"`
	Player string   `json:"player"`
}

type Pokemon struct {
	Name string `json:"name"`
}
type GetPokemonQuery struct {
	Player string `json:"player"`
}

type PokemonsPlayer struct {
	Pokemons []Pokemon
	Player   string
}

func CreatePokemonAddQuery(player string, names []string) (AddPokemonsQuery, error) {
	var retour AddPokemonsQuery
	if strings.TrimSpace(player) == "" {
		return retour, errors.New("player should not be empty")
	}
	return AddPokemonsQuery{Player: player, Names: names}, nil
}
