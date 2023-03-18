package controller

import (
	"clean/core"
	"gateway"
	"pokedex/domain"
	"presenter"
)

type PokedexController[T any] struct {
	ListPresenter     func() core.TransformPresenter[domain.PokemonsPlayer, T]
	GetPokemonGateway domain.IGetPokedemon
	AddPokemonGateway domain.IAddPokemon
}

func (this PokedexController[T]) GetMyPokemons(player string) core.IPresentOut[T] {
	useCase := domain.GetPokemonInPokedex{IGetPokedemon: this.GetPokemonGateway}
	presenter := this.ListPresenter()
	useCase.Execute(domain.GetPokemonQuery{Player: player}, &presenter)
	return presenter
}

func (this PokedexController[T]) AddPokemons(player string, names []string) core.IPresentOut[T] {

	presenter := this.ListPresenter()
	useCase := domain.AddPokemonInPokedex{IAddPokemon: this.AddPokemonGateway}
	query, err := domain.CreatePokemonAddQuery(player, names)

	if err != nil {
		presenter.ZeroValueErrorTransformePresenter(err)
	} else {
		useCase.Execute(query, &presenter)
	}

	return presenter
}

func NewControllerJSonAndMemory(repo gateway.Repo) PokedexController[string] {
	return PokedexController[string]{
		ListPresenter:     presenter.NewPokemonPlayerToJsonStringPresenter,
		GetPokemonGateway: gateway.GetAllMyPokemonGateway{Context: &repo},
		AddPokemonGateway: gateway.AddPokemonGateway{Context: &repo}}
}
