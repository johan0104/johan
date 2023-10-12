package hangman

import (
	"fmt"
	"strings"
)

func AffichageTirets(mot string) {
	motTirets := strings.Repeat("- ", len(mot))
	fmt.Println(mot)
	fmt.Println(motTirets)
}
