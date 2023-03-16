package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	AffichageMenu()
	ChoixMenu(reader)
}

func AffichageMenu() {
	fmt.Println("Hello  Bienvenu sur votre pokedex!!")
	fmt.Println("choisissez :")
	fmt.Println("1. Ajouter un pokemon")
}

func ChoixMenu(reader *bufio.Reader) string {
	ext, _ := reader.ReadString('\n')
	return ext
}
