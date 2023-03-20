package main

import (
	"bufio"
	"clean/core"
	"controller"
	"fmt"
	"gateway"
	"os"
	"pokedex/domain"
	"strings"
)

type configuration struct {
	Repo gateway.Repo
}

func main() {
	config := configuration{Repo: gateway.NewRepoForUnitTestsWithContext(map[string][]domain.Pokemon{"sacha": {{Name: "pikatchu"}}})}
	Start(config)
}

func Start(config configuration) {
	quit := false

	for !quit {
		reader := bufio.NewReader(os.Stdin)
		AffichageMenu()
		choix, _ := reader.ReadString('\n')
		switch strings.TrimSpace(choix) {
		case "q":
			quit = true
		case "1":
			ViewPokemon(config)
		default:
		}
	}
}

func AddPokemon() {
	fmt.Println("unimplemented")
}

func ViewPokemon(config configuration) {
	str, _ := GetPresenter(config).Print()
	fmt.Println(str)
}

func GetPresenter(config configuration) core.IPresentOut[string] {
	controller := controller.NewControllerForUnitTests(config.Repo)
	presenter := controller.GetMyPokemons("sacha")
	return presenter
}

func AffichageMenu() {
	fmt.Println("Hello  Bienvenu sur votre pokedex!!")
	fmt.Println("choisissez :")
	fmt.Println("1. Voir mes pokemons")
	fmt.Println(("q. Quit"))
}

func ChoixMenu(reader *bufio.Reader) string {
	ext, _ := reader.ReadString('\n')
	return ext
}
