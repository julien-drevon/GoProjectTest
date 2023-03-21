package gateway

import (
	"clean/core"
	"encoding/json"
	"io/ioutil"
	"pokedex/domain"
)

func NewRepoForUnitTests() Repo {
	return Repo{Context: make(map[string][]domain.Pokemon), PersistFile: false}
}

func NewRepoForWithPersistance(path string) (Repo, error) {
	var context map[string][]domain.Pokemon = make(map[string][]domain.Pokemon)
	var err error

	if core.IsExistFile(path) {
		context, err = UnserialyzeFile[map[string][]domain.Pokemon](path)

	}

	return Repo{Context: context, PersistFile: true, Path: path}, err
}

func UnserialyzeFile[T any](path string) (T, error) {
	//if core.IsExistFile(path) {
	var context T
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return context, err
	}
	err = json.Unmarshal(data, &context)

	return context, err

	//}
}
