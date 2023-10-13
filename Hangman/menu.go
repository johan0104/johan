package hangman

import "fmt"

func Menu() {
	var choixmenu int
	fmt.Println("1. Lancer le jeu ")
	fmt.Println("2. Score")
	fmt.Scan(&choixmenu)
	switch choixmenu {
	case 1:
		InterfaceJeu()
	case 2:
		Menuscore()
	default:
		fmt.Println("Valeur incorrecte")
		Menu()
	}
}
