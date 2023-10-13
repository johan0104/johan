package hangman

import "fmt"

func Menu() {
	var choixmenu int
	fmt.Println("1. Lancer le jeu ")
	fmt.Scan(&choixmenu)
	switch choixmenu {
	case 1:
		InterfaceJeu()
	default:
		fmt.Println("Valeur incorrecte")
		Menu()
	}
}
