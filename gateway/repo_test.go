package gateway

import (
	"clean/core"
	"os"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoSaveInFile(t *testing.T) {

	path := "TestRepoSaveInFilePokedex.json"
	assert := assert.New(t)
	DeleteRepoTestFile(path)

	repo, err := NewRepoForWithPersistance(path)
	assert.Equal(nil, err)

	err = repo.Add("sacha", domain.Pokemon{Name: "pikatchu"})
	assert.Equal(nil, err)

	expected := domain.PokemonsPlayer{Pokemons: []domain.Pokemon{{Name: "pikatchu"}}, Player: "sacha"}
	actual := repo.Get("sacha", func(x domain.Pokemon) bool { return true })
	assert.Equal(expected, actual)

	repo, err = NewRepoForWithPersistance(path)
	assert.Equal(nil, err)

	expected = domain.PokemonsPlayer{Pokemons: []domain.Pokemon{{Name: "pikatchu"}}, Player: "sacha"}
	actual = repo.Get("sacha", func(x domain.Pokemon) bool { return true })
	assert.Equal(expected, actual)

	DeleteRepoTestFile(path)
}

func DeleteRepoTestFile(path string) {

	if core.IsExistFile(path) {
		os.Remove(path)
	}
}
