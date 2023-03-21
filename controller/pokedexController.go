package controller

import (
	"clean/core"
	"gateway"
	"pokedex/domain"
	"presenter"
)

type PokedexController[T any] struct {
	ListPresenter      func() core.TransformPresenter[domain.PokemonsPlayer, T]
	GetPokemonGateway  domain.IGetPokedemon
	AddPokemonGateway  domain.IAddPokemon
	ReferentialGateway domain.IGetPokedex
}

func (this PokedexController[T]) GetMyPokemons(player string) core.IPresentOut[T] {
	useCase, presenter := this.InitGetMyPokemon()
	useCase.Execute(domain.GetPokemonQuery{Player: player}, &presenter)

	return presenter
}

func (this PokedexController[T]) AddPokemons(player string, names []string) core.IPresentOut[T] {

	presenter, useCase, query, err := this.InitAddPokemonFunction(player, names)

	if err != nil {
		presenter.ZeroValueErrorTransformPresenter(err)
	} else {
		useCase.Execute(query, &presenter)
	}

	return presenter
}

func (this PokedexController[T]) InitGetMyPokemon() (domain.GetPokemonInPokedex, core.TransformPresenter[domain.PokemonsPlayer, T]) {
	useCase := domain.GetPokemonInPokedex{IGetPokedemon: this.GetPokemonGateway}
	presenter := this.ListPresenter()
	return useCase, presenter
}

func (this PokedexController[T]) InitAddPokemonFunction(player string, names []string) (core.TransformPresenter[domain.PokemonsPlayer, T], domain.AddPokemonInPokedex, domain.AddPokemonsQuery, error) {
	presenter := this.ListPresenter()
	useCase := domain.AddPokemonInPokedex{IAddPokemon: this.AddPokemonGateway, IGetPokedex: this.ReferentialGateway}
	query, err := domain.CreatePokemonAddQuery(player, names)
	return presenter, useCase, query, err
}

func NewControllerForUnitTests(repo *gateway.Repo) PokedexController[string] {

	return PokedexController[string]{
		ListPresenter:      presenter.NewPokemonPlayerToJsonStringPresenter,
		GetPokemonGateway:  gateway.GetAllMyPokemonGateway{Context: repo},
		AddPokemonGateway:  gateway.AddPokemonGateway{Context: repo},
		ReferentialGateway: gateway.GetPokemonReferentialUnitTestsGateway{}}
}

func NewControllerWebMemory(repo *gateway.Repo) PokedexController[domain.PokemonsPlayer] {

	return PokedexController[domain.PokemonsPlayer]{
		ListPresenter:      presenter.NewPlayerPokemonWebServicePresenter,
		GetPokemonGateway:  gateway.GetAllMyPokemonGateway{Context: repo},
		AddPokemonGateway:  gateway.AddPokemonGateway{Context: repo},
		ReferentialGateway: gateway.GetPokemonReferentialFileGateway{}}
}
