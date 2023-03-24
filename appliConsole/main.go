package main

import (
	"bufio"
	"clean/core"
	"controller"
	"fmt"
	"gateway"
	"os"
	"strings"
)

type configuration struct {
	Repo gateway.Repo
}

func main() {
	repo, _ := gateway.NewRepoWithPersistance("pokedex1.json")
	config := &configuration{Repo: repo}
	Start(config)
}

func Start(config *configuration) {
	quit := false

	for !quit {
		reader := bufio.NewReader(os.Stdin)
		AffichageMenu()
		choix, _ := reader.ReadString('\n')
		switch strings.TrimSpace(choix) {
		case "q":
			quit = true

		case "1":
			ViewMyPokemon(config)

		case "2":
			ViewReferentiel()
		default:
		}
	}
}

func AddPokemon() {
	fmt.Println("unimplemented")
}

func ViewReferentiel() {
	str, _ := GetPokedexPresenterReferentiel().Print()
	fmt.Println(str)
}

func ViewMyPokemon(config *configuration) {
	str, _ := GetPokedexPresenter(config).Print()
	fmt.Println(str)
}

func GetPokedexPresenterReferentiel() core.IPresentOut[string] {
	controller := controller.NewPokemonReferentialForUnitsTests()
	presenter := controller.GetReferential()
	return presenter
}

func GetPokedexPresenter(config *configuration) core.IPresentOut[string] {
	controller := controller.NewControllerForUnitTests(&config.Repo)
	presenter := controller.GetMyPokemons("sacha")
	return presenter
}

func AffichageMenu() {
	fmt.Println("Hello  Bienvenu sur votre pokedex!!")
	fmt.Println("choisissez :")
	fmt.Println("1. Voir mes pokemons")
	fmt.Println("2. Voir tous les pokemons")
	fmt.Println(("q. Quit"))
}

func ChoixMenu(reader *bufio.Reader) string {
	ext, _ := reader.ReadString('\n')
	return ext
}
