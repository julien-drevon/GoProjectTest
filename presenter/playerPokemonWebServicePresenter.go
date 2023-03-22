package presenter

import (
	"clean/core"
	"net/http"
	"pokedex/domain"
)

type PlayerPokemonWebServicePresenter struct {
}

func (this PlayerPokemonWebServicePresenter) Convert(data domain.PokemonsPlayer, err error) (HttpResponse[domain.PokemonsPlayer], error) {
	if err != nil {
		var zeroV domain.PokemonsPlayer
		return HttpResponse[domain.PokemonsPlayer]{Data: zeroV, Status: http.StatusBadRequest, Error: err.Error()}, nil
	}

	return HttpResponse[domain.PokemonsPlayer]{Data: data, Status: http.StatusOK, Error: ""}, nil
}

func NewPlayerPokemonWebServicePresenter() core.TransformPresenter[domain.PokemonsPlayer, HttpResponse[domain.PokemonsPlayer]] {
	return core.TransformPresenter[domain.PokemonsPlayer, HttpResponse[domain.PokemonsPlayer]]{
		Converter: PlayerPokemonWebServicePresenter{}.Convert,
	}
}
