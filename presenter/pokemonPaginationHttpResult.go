package presenter

type PokemonPaginationHttpResult[T any] struct {
	Data   T
	Status int
	Error  string
}
