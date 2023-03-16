package pokemon

import (
	"clean/core"
	"encoding/json"
	"pokedex/domain"
)

type PokemonlistToJsonString struct {
}

func (converter PokemonlistToJsonString) Convert(data core.PaginationResult[domain.Pokemon]) (string, error) {
	pokeList := data.Result
	emp3, _ := json.Marshal(pokeList)
	return string(emp3), nil
}
