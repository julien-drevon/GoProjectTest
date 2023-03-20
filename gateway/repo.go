package gateway

import (
	"encoding/gob"
	"linq"
	"os"
	"pokedex/domain"
	"sync"
)

const pathOfPersistanceFile = "pokedexs.json"

type Repo struct {
	Context     map[string][]domain.Pokemon
	PersistFile bool
	Mutex       sync.Mutex
}

func (this *Repo) Get(name string, where func(domain.Pokemon) bool) domain.PokemonsPlayer {
	return domain.PokemonsPlayer{Player: name, Pokemons: linq.Where(this.Context[name], where)}
}

func (this *Repo) Add(playerName string, data domain.Pokemon) {
	this.Mutex.Lock()
	this.Context[playerName] = append(this.Context[playerName], data)
	if this.PersistFile {
		dataFile := this.OpenFile()
		this.SerializeFile(dataFile, data)
		dataFile.Close()
	}
	this.Mutex.Unlock()
}

func (*Repo) SerializeFile(dataFile *os.File, data domain.Pokemon) {
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(data)
}

func (*Repo) OpenFile() *os.File {
	dataFile, err := os.Create(pathOfPersistanceFile)
	if err != nil {
		os.Exit(1)
	}
	return dataFile
}
