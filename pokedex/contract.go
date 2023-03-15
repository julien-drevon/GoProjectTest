package domain

type IAddPokemon interface {
	Add(AddPokemonQuery) ([]Pokemon, error)
}
