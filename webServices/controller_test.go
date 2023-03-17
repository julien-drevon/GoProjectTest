package webservices

import (
	//"clean/core"
	// "pokedex/domain"
	// "presenter"

	"encoding/json"
	"linq"
	"net/http"
	"pokedex/domain"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func GetDebile() core.PaginationResult[domain.Pokemon] {
// 	return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
// }

type MyTestStruct struct {
	Id int
}

func TestCreateQuerySimple(t *testing.T) {
	assert := assert.New(t)

	query := domain.AddPokemonQuery{}
	expect := domain.AddPokemonQuery{Names: []string{"a"}}
	actual, _ := ConvertToQuery(query, func(q domain.AddPokemonQuery) (domain.AddPokemonQuery, error) {
		q.Names = []string{"a"}
		return q, nil
	})

	assert.Equal(expect, actual)
}

func TestCreateMultiQuery(t *testing.T) {
	assert := assert.New(t)

	query := domain.AddPokemonQuery{}
	expect := domain.AddPokemonQuery{Names: []string{"a1"}}
	actual, _ := MultiConvertToQuery(query, []func(domain.AddPokemonQuery) (domain.AddPokemonQuery, error){
		func(q domain.AddPokemonQuery) (domain.AddPokemonQuery, error) {
			q.Names = []string{"a"}
			return q, nil
		},
		func(q domain.AddPokemonQuery) (domain.AddPokemonQuery, error) {
			q.Names = linq.Select(q.Names, func(x string) string { return x + "1" })
			return q, nil
		}})

	assert.Equal(expect, actual)
}
func TestCreateQuerySimpleWithError(t *testing.T) {
	assert := assert.New(t)

	err, _ := json.Marshal(HTTPError{Code: http.StatusBadRequest, Message: ""})
	expect := string(err)
	_, actual := ConvertToQuery(MyTestStruct{}, func(q MyTestStruct) (MyTestStruct, error) {
		id, internalErr := strconv.Atoi("a")
		q.Id = id
		return q, internalErr
	})

	assert.Equal(expect, actual.Error())
}
