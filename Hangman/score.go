package hangman

import "fmt"

func Menuscore() {
	var choixmenuscore int
	fmt.Println("Votre score est de ...")
	fmt.Println("0. Quitter")
	fmt.Scan(&choixmenuscore)
	switch choixmenuscore {
	case 0:
		Menu()
	default:
		fmt.Println("Valeur incorrecte")
		Menuscore()
	}
}
