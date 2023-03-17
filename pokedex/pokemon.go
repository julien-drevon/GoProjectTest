package domain

// type IExecuteUseCase[TQuery any, TData any] interface {
// 	Execute(TQuery, IPresentIn[TData])
// }
type AddPokemonQuery struct {
	Names []string `json:"names"`
}

type Pokemon struct {
	Name string `json:"name"`
}
type GetPokemonQuery struct {
	Name string `json:"name"`
}
