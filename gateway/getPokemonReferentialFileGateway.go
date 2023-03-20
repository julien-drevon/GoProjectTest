package gateway

import (
	"clean/core"
	"encoding/json"
	"errors"
	"io/ioutil"
	"linq"
	"os"
	"pokedex/domain"
)

type GetPokemonReferentialFleGateway struct {
}

var cache []domain.Pokemon

func (this GetPokemonReferentialFleGateway) GetPokedex(query domain.GetPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	val, err := getCache()
	return core.NewPaginationResult(val, len(val), 1, 0), err
}

func getCache() ([]domain.Pokemon, error) {

	if cache != nil {
		return cache, nil
	}

	zeroValue := []domain.Pokemon{}
	jsonFile, errOpenFile := openFileList()

	if errOpenFile != nil {
		return zeroValue, errOpenFile
	}

	pokeLi, errUn := Unserialyze(jsonFile)
	if errUn != nil {
		return zeroValue, errUn
	}

	if jsonFile != nil {
		defer jsonFile.Close()
	}

	cache = pokeLi
	return cache, nil
}

func Unserialyze(fileStream *os.File) ([]domain.Pokemon, error) {
	byteValue, _ := ioutil.ReadAll(fileStream)
	var result []domain.Pokemon

	err := json.Unmarshal(byteValue, &result)
	if err != nil {
		err = errors.Join(errors.New("Unmarshall error. "), err)
	}
	return result, err
}

func openFileList() (*os.File, error) {
	jsonFile, err := os.Open("pokemons.json")

	if err != nil {
		var defaultFileValue *os.File
		err = errors.Join(errors.New("json file opening errror. "), err)
		return defaultFileValue, err
	}

	return jsonFile, nil
}

func (this GetPokemonReferentialFleGateway) IsExist(query domain.AddPokemonsQuery) bool {
	return linq.Any(cache, func(x domain.Pokemon) bool { return linq.Any(query.Names, func(s string) bool { return s == x.Name }) })
}
