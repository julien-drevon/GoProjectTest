package gateway

import (
	"clean/core"
	"linq"
	"pokedex/domain"
)

var POKEMONREFERENTIEL_CACHE []domain.Pokemon

const PATH_POKEMONREFERENTIEL_FILE = "pokemons.json"

type GetPokemonReferentialFileGateway struct {
}

func (this GetPokemonReferentialFileGateway) GetPokedex(query domain.GetPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	val, err := GetCache()
	return core.NewPaginationResult(val, len(val), 1, 0), err
}

func GetCache() ([]domain.Pokemon, error) {
	path := PATH_POKEMONREFERENTIEL_FILE
	if POKEMONREFERENTIEL_CACHE == nil {
		zeroValue := []domain.Pokemon{}
		pokeLi, errUn := core.UnserialyzeFile[[]domain.Pokemon](path)
		if errUn != nil {
			return zeroValue, errUn
		}
		POKEMONREFERENTIEL_CACHE = pokeLi
	}
	return POKEMONREFERENTIEL_CACHE, nil
}

func (this GetPokemonReferentialFileGateway) IsExist(query domain.AddPokemonsQuery) bool {
	return linq.Any(POKEMONREFERENTIEL_CACHE, func(x domain.Pokemon) bool { return linq.Any(query.Names, func(s string) bool { return s == x.Name }) })
}
