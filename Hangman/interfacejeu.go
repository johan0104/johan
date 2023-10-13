package hangman

import (
	"fmt"
	"strings"
)

func InterfaceJeu() {
	SupprimerMot()
	Ecriremot()
	fmt.Println("Le mot à deviner : ")
	AffichageTirets("Salut")
	tentatives := 10

	var MotTrouvé bool
	var choixlettres string
	for tentatives > 0 && MotTrouvé == false {
		fmt.Printf("Il vous reste : %d tentatives\n", tentatives)
		fmt.Println("Entrez une lettre ou le mot entier : ")
		fmt.Scan(&choixlettres)
		if len(choixlettres) == 1 {
			if strings.Contains("salut", choixlettres) {
				fmt.Println("Vous avez trouvé une lettre")
			} else {
				fmt.Println("Cette lettre n'est pas dans le mot")
				tentatives--
			}
		}
		if len(choixlettres) > 1 {
			if strings.Contains("salut", choixlettres) {
				fmt.Println("Vous avez trouvé le mot")
				MotTrouvé = true
			} else {
				fmt.Println("Le mot n'est pas bon")
				tentatives--
			}
		}

	}
	if MotTrouvé == true {
		fmt.Println("Vous avez gagné")
		Menu()
	} else {
		fmt.Println("Vous n'avez plus de tentatives")
		Menu()
	}
}
