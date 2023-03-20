package gateway

import (
	"os"
	"pokedex/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoSaveInFile(t *testing.T) {
	assert := assert.New(t)
	DeleteFile()
	repo := NewRepoForWithPersistance()

	repo.Add("sacha", domain.Pokemon{Name: "pikatchu"})

	expected := domain.PokemonsPlayer{Pokemons: []domain.Pokemon{{Name: "pikatchu"}}, Player: "sacha"}
	actual := repo.Get("sacha", func(x domain.Pokemon) bool { return true })
	assert.Equal(expected, actual)

	repo = NewRepoForWithPersistance()
	expected = domain.PokemonsPlayer{Pokemons: []domain.Pokemon{{Name: "pikatchu"}}, Player: "sacha"}
	actual = repo.Get("sacha", func(x domain.Pokemon) bool { return true })
	assert.Equal(expected, actual)

	DeleteFile()
}

func DeleteFile() {
	os.Remove(pathOfPersistanceFile)
}
