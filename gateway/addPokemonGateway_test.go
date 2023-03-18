package gateway

import (
	//"clean/core"
	// "pokedex/domain"
	// "presenter"

	"clean/core"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func GetDebile() core.PaginationResult[domain.Pokemon] {
// 	return core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
// }

func Test2AddPokemonMemory(t *testing.T) {
	assert := assert.New(t)
	repo := NewRepo[domain.Pokemon]()
	service := AddPokemonGateway{Context: &repo}
	expected := core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
	actual, _ := service.Add(domain.AddPokemonQuery{Names: []string{"pikatchu", "tortank"}})

	assert.Equal(expected, actual)
}
func Test2AddPokemonMemoryWithOther(t *testing.T) {
	assert := assert.New(t)
	repo := NewRepoWithContext([]domain.Pokemon{{Name: "draco feu"}})
	service := AddPokemonGateway{Context: &repo}
	service.Add(domain.AddPokemonQuery{Names: []string{"pikatchu", "tortank"}})

	expected := []domain.Pokemon{{Name: "draco feu"}, {Name: "pikatchu"}, {Name: "tortank"}}
	actual := service.Context.Get()

	assert.Equal(expected, actual)
}
