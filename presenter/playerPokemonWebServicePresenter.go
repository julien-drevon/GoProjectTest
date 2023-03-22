package presenter

import (
	"clean/core"
	"pokedex/domain"
)

type PlayerPokemonWebServicePresenter struct {
}

func (this PlayerPokemonWebServicePresenter) Convert(data domain.PokemonsPlayer, err error) (domain.PokemonsPlayer, error) {
	return data, err
}

func NewPlayerPokemonWebServicePresenter() core.TransformPresenter[domain.PokemonsPlayer, domain.PokemonsPlayer] {
	return core.TransformPresenter[domain.PokemonsPlayer, domain.PokemonsPlayer]{
		Converter: PlayerPokemonWebServicePresenter{}.Convert,
	}
}
