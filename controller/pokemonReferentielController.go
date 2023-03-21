package controller

import (
	"clean/core"
	"pokedex/domain"
)

type PokemonReferentialController[T any] struct {
	ReferentialPresenter func() core.TransformPresenter[core.PaginationResult[domain.Pokemon], T]
	ReferentialGateway   domain.IGetPokedex
}

func (this PokemonReferentialController[T]) GetReferential() core.IPresentOut[T] {
	presenter, useCase, query := this.InitGetReferential()
	useCase.Execute(query, &presenter)
	return presenter
}

func (this PokemonReferentialController[T]) InitGetReferential() (core.TransformPresenter[core.PaginationResult[domain.Pokemon], T], domain.GetPokemonReferential, domain.GetPokemonQuery) {
	presenter := this.ReferentialPresenter()
	useCase := domain.GetPokemonReferential{IGetPokedex: this.ReferentialGateway}
	query := domain.GetPokemonQuery{}
	return presenter, useCase, query
}
