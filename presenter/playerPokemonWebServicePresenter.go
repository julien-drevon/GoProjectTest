package presenter

import (
	"clean/core"
	"pokedex/domain"
)

type PlayerPokemonWebServicePresenter struct {
}

func (converter PlayerPokemonWebServicePresenter) Convert(data domain.PokemonsPlayer) (domain.PokemonsPlayer, error) {
	return data, nil
}

func NewPlayerPokemonWebServicePresenter() core.TransformPresenter[domain.PokemonsPlayer, domain.PokemonsPlayer] {
	return core.TransformPresenter[domain.PokemonsPlayer, domain.PokemonsPlayer]{
		Converter: PlayerPokemonWebServicePresenter{}.Convert,
	}
}
