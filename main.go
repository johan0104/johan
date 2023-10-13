package main

import (
	"fmt"
	h "hangman/Hangman"
)

func main() {
	randomWord, err := h.GetRandomWordFromFile("Hangman/dico.txt")
	if err != nil {
		fmt.Println("Erreur de lecture:", err)
		return
	}

	fmt.Println("Le mot choisi est:", randomWord)
}
