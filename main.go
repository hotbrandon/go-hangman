package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func OpenDictionary(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// count the number of words whose length is between minLen and maxLen
func GetMatchingWords(file *os.File, minLen, maxLen int) []string {
	scanner := bufio.NewScanner(file)
	var words []string

	for scanner.Scan() {
		// Scanner.Text() automatically strips the newline character (\n),
		line := scanner.Text()
		if len(line) >= minLen && len(line) <= maxLen {
			words = append(words, line)
		}
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
		return nil
	}

	return words
}

func GetRandomWord(words []string) string {
	return words[rand.IntN(len(words))]
}

func GetInitialWord(word string) string {
	runes := []rune(word)
	length := len(runes)
	if length <= 2 {
		return word
	}

	for i := 1; i < length-1; i++ {
		runes[i] = '_'
	}

	return string(runes)
}

func CheckValidAlphabet(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func main() {
	file, err := OpenDictionary("words_alpha.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {
		chosenWord := GetRandomWord(GetMatchingWords(file, 3, 10))
		max_guess_count := len(chosenWord) * 2
		guessed_word := GetInitialWord(chosenWord)
		for i := 0; i < max_guess_count; i++ {
			fmt.Printf("word: %s\n", guessed_word)
			fmt.Println("Guess the next letter:")
			reader := bufio.NewReader(os.Stdin)

			input, _ := reader.ReadString('\n')
			letter := strings.TrimSpace(input)

			fmt.Println(letter)
		}
	}

}
