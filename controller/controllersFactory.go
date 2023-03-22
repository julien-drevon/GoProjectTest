package controller

import (
	"clean/core"
	"gateway"
	"pokedex/domain"
	"presenter"
)

func NewControllerForUnitTests(repo *gateway.Repo) PokedexController[string] {
	return PokedexController[string]{
		ListPresenter:      presenter.NewPokemonPlayerToJsonStringPresenter,
		GetPokemonGateway:  gateway.GetAllMyPokemonGateway{Context: repo},
		AddPokemonGateway:  gateway.AddPokemonGateway{Context: repo},
		ReferentialGateway: gateway.GetPokemonReferentialUnitTestsGateway{}}
}

func NewControllerFileWeb(repo *gateway.Repo) PokedexController[domain.PokemonsPlayer] {
	return PokedexController[domain.PokemonsPlayer]{
		ListPresenter:      presenter.NewPlayerPokemonWebServicePresenter,
		GetPokemonGateway:  gateway.GetAllMyPokemonGateway{Context: repo},
		AddPokemonGateway:  gateway.AddPokemonGateway{Context: repo},
		ReferentialGateway: gateway.GetPokemonReferentialFileGateway{}}
}

func NewPokemonReferentialForUnitsTests() PokemonReferentialController[string] {

	return PokemonReferentialController[string]{
		ReferentialPresenter: presenter.NewPokemonListToJsonStringPresenter,
		ReferentialGateway:   gateway.GetPokemonReferentialUnitTestsGateway{}}
}

func NewReferentialController() PokemonReferentialController[presenter.HttpResponse[core.PaginationResult[domain.Pokemon]]] {

	return PokemonReferentialController[presenter.HttpResponse[core.PaginationResult[domain.Pokemon]]]{
		ReferentialPresenter: presenter.NewPokemonPaginationWebServicePresenter,
		ReferentialGateway:   gateway.GetPokemonReferentialFileGateway{}}
}
