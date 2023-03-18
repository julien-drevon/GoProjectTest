package domain

type IAddPokemon interface {
	Add(AddPokemonsQuery) (PokemonsPlayer, error)
}
type IGetPokedemon interface {
	Get(GetPokemonQuery) (PokemonsPlayer, error)
}
type IGetPokedex interface {
	GetPokedex(GetPokemonQuery) (PokemonsPlayer, error)
}
