package main

import (
	"bufio"
	"clean/core"
	"fmt"
	"os"
	"pokedex/domain"
	"presenter"
	"strings"
)

func main() {

	quit := false

	for !quit {
		reader := bufio.NewReader(os.Stdin)
		AffichageMenu()
		choix, _ := reader.ReadString('\n')
		switch strings.TrimSpace(choix) {
		case "q":
			quit = true
		case "1":
			ViewPokemon()
		default:
		}

	}
}
func AddPokemon() {
	fmt.Println("unimplemented")
}

func ViewPokemon() {
	pres, _ := GetPresenter()
	fmt.Println(pres.Print())
}

func GetPresenter() (core.IPresentOut[string], error) {
	pres := &core.TransformPresenter[core.PaginationResult[domain.Pokemon], string]{Converter: presenter.PokemonListToJsonStringConverter{}}
	li := core.NewPaginationResult([]domain.Pokemon{{Name: "pikatchu"}, {Name: "tortank"}}, 2, 1, 0)
	pres.Present(li, nil)
	return pres, nil
}

func AffichageMenu() {
	fmt.Println("Hello  Bienvenu sur votre pokedex!!")
	fmt.Println("choisissez :")
	fmt.Println("1. Ajouter un pokemon")
	fmt.Println(("q. Quit"))
}

func ChoixMenu(reader *bufio.Reader) string {
	ext, _ := reader.ReadString('\n')
	return ext
}
