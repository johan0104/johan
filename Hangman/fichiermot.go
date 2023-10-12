package hangman

import (
	"os"
)

func Ecriremot() {
	mot := []byte("Salut")
	os.WriteFile("mot.txt", mot, 0644)

}
