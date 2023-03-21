package gateway

import (
	"clean/core"
	"pokedex/domain"
)

func NewRepoForUnitTests() Repo {
	return Repo{Context: make(map[string][]domain.Pokemon), PersistFile: false}
}

func NewRepoForWithPersistance(path string) (Repo, error) {
	var context map[string][]domain.Pokemon = make(map[string][]domain.Pokemon)
	var err error

	if core.IsExistFile(path) {
		context, err = core.UnserializeFile[map[string][]domain.Pokemon](path)
	}

	return Repo{Context: context, PersistFile: true, Path: path}, err
}
