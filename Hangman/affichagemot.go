package hangman

import (
	"fmt"
)

// Affiche le mot choisi en tiret avec un espace entre chaque tiret
func AffichageTirets(Word string) {
	for _, lettre := range Word{
		lettreTrouvee := false
		for _, lettresTrouvee := range LettresTrouvees {
			if string(lettre) == LettresTrouvees {
				lettresTrouvee = true
				break
			}
		}
		if lettreTrouvee {
			fmt.Print(string(lettre), " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}
