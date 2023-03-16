package main

import (
	"bufio"
	"clean/core"
	"controller"
	"fmt"
	"os"
	"pokedex/domain"
	"strings"
)

type configuration struct {
	pokemonLi []domain.Pokemon
}

func main() {

	quit := false
	config := configuration{pokemonLi: make([]domain.Pokemon, 0)}

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
	controller := controller.NewControllerJSonAndMemory(config.pokemonLi)
	presenter := controller.GetMyPokemons()
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
