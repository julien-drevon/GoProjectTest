package domain

type IAddPokemon interface {
	Add(AddPokemonQuery) ([]Pokemon, error)
}
type IGetPokemon interface {
	Get(GetPokemonQuery) ([]Pokemon, error)
}
