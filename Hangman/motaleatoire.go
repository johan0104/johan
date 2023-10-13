package hangman

import (
	"bufio"
	"os"
	"time"
)

func GetRandomWordFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	randomIndex := GenerateRandomNumber(len(words))

	return words[randomIndex], nil
}

func GenerateRandomNumber(max int) int {
	seed := time.Now().UnixNano()
	random := seed % int64(max)
	return int(random)
}

