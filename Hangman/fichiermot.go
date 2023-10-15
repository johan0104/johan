package hangman

import (
	"fmt"
	"os"
)

func Ecriremot() {
	filename := "Hangman/dico.txt"
	word, err := GetRandomWordFromFile(filename)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	os.WriteFile("mot.txt", []byte(word), 0644)
}

func SupprimerMot() {
	os.Remove("mot.txt")
}
