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

// return a slice of words whose length is between minLen and maxLen from the input file
func FilterWordsByLength(file *os.File, minLen, maxLen int) []string {
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

// display the first and last alphabet of a word, and hide others with a '_'
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

func findLetterPosition(letter rune, word string) []int {
	var indexes []int
	for idx, char := range word {
		if letter == char {
			indexes = append(indexes, idx)
			fmt.Printf("found at index %d, letter %c\n", idx, char)
		}
	}
	return indexes
}

func fillinLetter(guessedWord string, letter rune, positions []int) string {
	runes := []rune(guessedWord)
	for _, pos := range positions {
		runes[pos] = letter
	}
	return string(runes)
}

func main() {
	file, err := OpenDictionary("words_alpha.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	word_list := FilterWordsByLength(file, 3, 10)
	for {
		chosenWord := GetRandomWord(word_list)
		maxGuessCount := len(chosenWord)
		fmt.Printf("max try %d, enter a letter or 'quit' to exit.\n", maxGuessCount)
		guessed_word := GetInitialWord(chosenWord)
		guessed_letters := make(map[rune]bool)

		for i := 0; i < maxGuessCount; i++ {
			fmt.Printf("word: %s\n", guessed_word)
			fmt.Printf("you have %d tries left.\n", maxGuessCount-i)
			fmt.Println("Guess the next letter:")
			reader := bufio.NewReader(os.Stdin)

			input, _ := reader.ReadString('\n')
			trimmedInput := strings.TrimSpace(input)

			if trimmedInput == "quit" {
				fmt.Println("good bye!")
				os.Exit(0)
			}
			if len(trimmedInput) == 0 {
				fmt.Println("invalid alphabet, please use 'A' - 'Z' or 'a' - 'z'")
				continue
			}

			if !CheckValidAlphabet(rune(trimmedInput[0])) {
				fmt.Println("invalid alphabet, please use 'A' - 'Z' or 'a' - 'z'")
				continue
			}
			guessed_letter := rune(trimmedInput[0])
			if guessed_letters[guessed_letter] {
				fmt.Printf("you have already guessed the letter %c.\n", guessed_letter)
			}
			guessed_letters[guessed_letter] = true

			// find the letter in chosenWord, and return the position(s).
			positions := findLetterPosition(rune(trimmedInput[0]), chosenWord)
			guessed_word = fillinLetter(guessed_word, rune(trimmedInput[0]), positions)

			if guessed_word == chosenWord {
				fmt.Printf("Congratulations! You guessed the word: %s\n\n", chosenWord)
				break
			}

			fmt.Println("guessed letters: ")
			for c := range guessed_letters {
				fmt.Printf("%c ", c)
			}
			fmt.Println()
			if i == maxGuessCount-1 {
				fmt.Printf("Game over! The word was: %s\n\n", chosenWord)
			}
		}
	}

}
