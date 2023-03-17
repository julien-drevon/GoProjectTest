package controller

import (
	"clean/core"
	"gateway"
	"pokedex/domain"
	"presenter"
)

type PokedexController[T any] struct {
	ListPresenter     func() core.TransformPresenter[core.PaginationResult[domain.Pokemon], T]
	GetPokemonGateway domain.IGetPokedemon
	AddPokemonGateway domain.IAddPokemon
}

func (this PokedexController[T]) GetMyPokemons() core.IPresentOut[T] {
	useCase := domain.GetPokemonInPokedex{IGetPokedemon: this.GetPokemonGateway}
	presenter := this.ListPresenter()
	useCase.Execute(domain.GetPokemonQuery{}, &presenter)
	return presenter
}

func (this PokedexController[T]) AddPokemons(name string) core.IPresentOut[T] {
	useCase := domain.AddPokemonInPokedex{IAddPokemon: this.AddPokemonGateway}
	presenter := this.ListPresenter()
	useCase.Execute(domain.AddPokemonQuery{Names: []string{name}}, &presenter)
	return presenter
}

func NewControllerJSonAndMemory(pokemonLi []domain.Pokemon) PokedexController[string] {
	return PokedexController[string]{ListPresenter: presenter.NewPokemonListToJsonStringPresenter, GetPokemonGateway: gateway.GetAllMyPokemonGateway{PokeList: pokemonLi}}
}