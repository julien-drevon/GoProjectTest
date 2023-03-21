package gateway

import (
	"clean/core"
	"linq"
	"pokedex/domain"
	"sync"
)

//const pathOfPersistanceFile = "pokedexs.json"

type Repo struct {
	Context     map[string][]domain.Pokemon
	PersistFile bool
	Mutex       sync.Mutex
	Path        string
}

func (this *Repo) Get(name string, where func(domain.Pokemon) bool) domain.PokemonsPlayer {
	return domain.PokemonsPlayer{Player: name, Pokemons: linq.Where(this.Context[name], where)}
}

func (this *Repo) Add(playerName string, data domain.Pokemon) error {

	var err error

	this.Mutex.Lock()
	this.Context[playerName] = append(this.Context[playerName], data)

	if this.PersistFile {
		err = core.SerializeFile(this.Path, this.Context)
	}

	this.Mutex.Unlock()
	return err
}
