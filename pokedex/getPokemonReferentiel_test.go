package domain

// import (
// 	"clean/core"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// type GetAllPokemonsReferentialStub struct {
// }

// func (service GetAllPokemonsReferentialStub) GetPokedex(query GetPokemonQuery) (core.PaginationResult[Pokemon], error) {
// 	return core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 4, 1, 0), nil
// }

// func Test_GetAllPokemonsReferential(t *testing.T) {
// 	assert := assert.New(t)

// 	query := GetPokemonQuery{Player: "sacha"}
// 	useCase := GetPokemonReferential{IGetPokedex: GetAllPokemonsReferentialStub{}}
// 	presenter := &core.SimplePresenter[core.PaginationResult[Pokemon]]{}
// 	useCase.Execute(query, presenter)

// 	actual, _ := presenter.Print()
// 	expected := core.NewPaginationResult([]Pokemon{{Name: "pikatchu"}, {Name: "Tortank"}}, 4, 1, 0)

// 	assert.Equal(expected, actual)
// }
