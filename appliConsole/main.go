package main

import (
	"bufio"
	"fmt"
	"os"
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
			AddPokemon()
		default:
		}

	}
}
func AddPokemon() {
	fmt.Println("unimplemented")
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
