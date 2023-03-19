package controller

import (
	"clean/core"
	"gateway"
	"pokedex/domain"
	"presenter"
)

type PokeController[T any] struct {
	ListPresenter        func() core.TransformPresenter[domain.PokemonsPlayer, T]
	ReferentielPresenter func() core.TransformPresenter[core.PaginationResult[domain.Pokemon], T]
	GetPokemonGateway    domain.IGetPokedemon
	AddPokemonGateway    domain.IAddPokemon
	ReferentielGateway   domain.IGetPokedex
}

func (this PokeController[T]) GetMyPokemons(player string) core.IPresentOut[T] {
	useCase, presenter := this.InitGetMyPokemon()
	useCase.Execute(domain.GetPokemonQuery{Player: player}, &presenter)

	return presenter
}

func (this PokeController[T]) AddPokemons(player string, names []string) core.IPresentOut[T] {

	presenter, useCase, query, err := this.InitAddPokemonFunction(player, names)

	if err != nil {
		presenter.ZeroValueErrorTransformePresenter(err)
	} else {
		useCase.Execute(query, &presenter)
	}

	return presenter
}

func (this PokeController[T]) GetReferentiel() core.IPresentOut[T] {

	presenter, useCase, query := this.InitGetReferentiel()
	useCase.Execute(query, &presenter)

	return presenter
}

func (this PokeController[T]) InitGetMyPokemon() (domain.GetPokemonInPokedex, core.TransformPresenter[domain.PokemonsPlayer, T]) {
	useCase := domain.GetPokemonInPokedex{IGetPokedemon: this.GetPokemonGateway}
	presenter := this.ListPresenter()
	return useCase, presenter
}

func (this PokeController[T]) InitAddPokemonFunction(player string, names []string) (core.TransformPresenter[domain.PokemonsPlayer, T], domain.AddPokemonInPokedex, domain.AddPokemonsQuery, error) {
	presenter := this.ListPresenter()
	useCase := domain.AddPokemonInPokedex{IAddPokemon: this.AddPokemonGateway}
	query, err := domain.CreatePokemonAddQuery(player, names)
	return presenter, useCase, query, err
}

func (this PokeController[T]) InitGetReferentiel() (core.TransformPresenter[core.PaginationResult[domain.Pokemon], T], domain.GetPokemonReferentiel, domain.GetPokemonQuery) {
	presenter := this.ReferentielPresenter()
	useCase := domain.GetPokemonReferentiel{IGetPokedex: this.ReferentielGateway}
	query := domain.GetPokemonQuery{}
	return presenter, useCase, query
}

func NewControllerJSonAndMemory(repo gateway.Repo) PokeController[string] {

	return PokeController[string]{
		ListPresenter:        presenter.NewPokemonPlayerToJsonStringPresenter,
		ReferentielPresenter: presenter.NewPokemonListToJsonStringPresenter,
		GetPokemonGateway:    gateway.GetAllMyPokemonGateway{Context: &repo},
		AddPokemonGateway:    gateway.AddPokemonGateway{Context: &repo},
		ReferentielGateway:   gateway.GetPokemonReferentielGateway{}}
}
