package gateway

import (
	"linq"
	"pokedex/domain"
)

type Repo struct {
	Context map[string][]domain.Pokemon
}

func (this *Repo) Get(name string, where func(domain.Pokemon) bool) domain.PokemonsPlayer {
	return domain.PokemonsPlayer{Player: name, Pokemons: linq.Where(this.Context[name], where)}
}

func (this *Repo) Add(playerName string, data domain.Pokemon) {
	this.Context[playerName] = append(this.Context[playerName], data)
}

func NewRepo() Repo {
	return Repo{Context: make(map[string][]domain.Pokemon)}
}

func NewRepoWithContext(context map[string][]domain.Pokemon) Repo {
	return Repo{Context: context}
}
