package presenter

import (
	"clean/core"
	"net/http"
	"pokedex/domain"
)

type PokemonPaginationWebServicePresenter struct {
}

func (this PokemonPaginationWebServicePresenter) Convert(data core.PaginationResult[domain.Pokemon], err error) (HttpResponse[core.PaginationResult[domain.Pokemon]], error) {
	if err != nil {
		return HttpResponse[core.PaginationResult[domain.Pokemon]]{Status: http.StatusBadRequest, Error: err.Error(), Data: core.PaginationResult[domain.Pokemon]{}}, nil
	}

	return HttpResponse[core.PaginationResult[domain.Pokemon]]{Status: http.StatusOK, Data: data, Error: ""}, nil
}

// domain.PokemonsPlayer
func NewPokemonPaginationWebServicePresenter() core.TransformPresenter[core.PaginationResult[domain.Pokemon], HttpResponse[core.PaginationResult[domain.Pokemon]]] {
	return core.TransformPresenter[core.PaginationResult[domain.Pokemon], HttpResponse[core.PaginationResult[domain.Pokemon]]]{
		Converter: PokemonPaginationWebServicePresenter{}.Convert,
	}
}
