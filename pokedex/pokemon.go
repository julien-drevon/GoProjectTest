package domain

// type IExecuteUseCase[TQuery any, TData any] interface {
// 	Execute(TQuery, IPresentIn[TData])
// }
type AddPokemonsQuery struct {
	Names  []string `json:"names"`
	Player string   `json:"player"`
}

type Pokemon struct {
	Name string `json:"name"`
}
type GetPokemonQuery struct {
	Player string `json:"player"`
}

type PokemonsPlayer struct {
	Pokemons []Pokemon
	Player   string
}
