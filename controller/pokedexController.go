package controller

import (
	"clean/core"
	"pokedex/domain"
)

type PokedexController[T any] struct {
	Presenter core.TransformPresenter[core.PaginationResult[domain.Pokemon], T]
	Gateway   domain.IGetPokedemon
}

func (this PokedexController[T]) GetMyPokemons() core.IPresentOut[T] {
	useCase := domain.GetPokemonInPokedex{IGetPokedemon: this.Gateway}
	useCase.Execute(domain.GetPokemonQuery{}, &this.Presenter)
	return this.Presenter
}
