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
	tentatives := 1
	maxtentatives := 10
	var choixlettres string
	for tentatives <= maxtentatives {
		fmt.Printf("Il vous reste : %d tentatives\n", tentatives)
		fmt.Println("Entrez une lettre ou le mot entier : ")
		fmt.Scan(&choixlettres)
		if len(choixlettres) == 1 {
			if strings.Contains("salut", choixlettres) {
				fmt.Println("Vous avez trouvé une lettre")
			} else {
				fmt.Println("Cette lettre n'est pas dans le mot")
				tentatives++
			}
		}

	}
	fmt.Println("Vous n'avez plus de tentatives")
}
