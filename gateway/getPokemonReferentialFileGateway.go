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

// var POKEDEX []domain.Pokemon = []domain.Pokemon{
// 	{Name: "draco feu"},
// 	{Name: "pikatchu"},
// 	{Name: "tortank"},
// }

func (this GetPokemonReferentialFleGateway) GetPokedex(query domain.GetPokemonQuery) (core.PaginationResult[domain.Pokemon], error) {
	zeroValue := core.NewPaginationResult([]domain.Pokemon{}, 0, 1, 0)

	jsonFile, errOpenFile := this.openFileList()
	if errOpenFile != nil {
		return zeroValue, errOpenFile
	}

	pokeLi, errUn := this.Unserialyze(jsonFile)
	if errUn != nil {
		return zeroValue, errUn
	}
	if jsonFile != nil {
		defer jsonFile.Close()
	}

	return core.NewPaginationResult(pokeLi, len(pokeLi), 1, 0), nil
}

func (this GetPokemonReferentialFleGateway) Unserialyze(fileStream *os.File) ([]domain.Pokemon, error) {
	byteValue, _ := ioutil.ReadAll(fileStream)
	var result []domain.Pokemon

	err := json.Unmarshal(byteValue, &result)
	if err != nil {
		err = errors.Join(errors.New("Unmarshall error. "), err)
	}
	return result, err
}

func (this GetPokemonReferentialFleGateway) openFileList() (*os.File, error) {
	jsonFile, err := os.Open("pokemons.json")

	if err != nil {
		var defaultFileValue *os.File
		err = errors.Join(errors.New("json file opening errror. "), err)
		return defaultFileValue, err
	}
	//defer jsonFile.Close()
	return jsonFile, nil
}

func (this GetPokemonReferentialFleGateway) IsExist(query domain.AddPokemonsQuery) bool {
	return linq.Any(POKEDEX, func(x domain.Pokemon) bool { return linq.Any(query.Names, func(s string) bool { return s == x.Name }) })
}
