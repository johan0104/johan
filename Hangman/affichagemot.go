package hangman

import (
	"fmt"
	"strings"
)

// Affiche le mot choisi en tiret avec un espace entre chaque tiret
func AffichageTirets(ContenueFichierMot string) {
	motTirets := strings.Repeat("- ", len(ContenueFichierMot))
	fmt.Println(motTirets)
}
