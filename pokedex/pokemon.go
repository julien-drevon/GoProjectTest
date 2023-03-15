package domain

// type IExecuteUseCase[TQuery any, TData any] interface {
// 	Execute(TQuery, IPresentIn[TData])
// }
type AddPokemonQuery struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name string `json:"name"`
}
