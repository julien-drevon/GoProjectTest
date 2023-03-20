package controller

import (
	"clean/core"
	"gateway"
	"pokedex/domain"
	"presenter"
)

type PokemonReferentielController[T any] struct {
	ReferentielPresenter func() core.TransformPresenter[core.PaginationResult[domain.Pokemon], T]
	ReferentielGateway   domain.IGetPokedex
}

func (this PokemonReferentielController[T]) GetReferentiel() core.IPresentOut[T] {

	presenter, useCase, query := this.InitGetReferentiel()
	useCase.Execute(query, &presenter)

	return presenter
}

func (this PokemonReferentielController[T]) InitGetReferentiel() (core.TransformPresenter[core.PaginationResult[domain.Pokemon], T], domain.GetPokemonReferentiel, domain.GetPokemonQuery) {
	presenter := this.ReferentielPresenter()
	useCase := domain.GetPokemonReferentiel{IGetPokedex: this.ReferentielGateway}
	query := domain.GetPokemonQuery{}
	return presenter, useCase, query
}

func NewPokemonReferentielControllerJSonAndMemory() PokemonReferentielController[string] {

	return PokemonReferentielController[string]{
		ReferentielPresenter: presenter.NewPokemonListToJsonStringPresenter,
		ReferentielGateway:   gateway.GetPokemonReferentielGateway{}}
}

func NewReferentielController() PokemonReferentielController[core.PaginationResult[domain.Pokemon]] {

	return PokemonReferentielController[core.PaginationResult[domain.Pokemon]]{

		ReferentielPresenter: presenter.NewPokemonPaginationWebServicePresenter,
		ReferentielGateway:   gateway.GetPokemonReferentielGateway{}}
}
