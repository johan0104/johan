package hangman

import (
	"fmt"
	"strings"
)

// Affiche le mot choisi en tiret avec un espace entre chaque tiret
func AffichageTirets(mot string) {
	motTirets := strings.Repeat("- ", len(mot))
	fmt.Println(motTirets)
}
