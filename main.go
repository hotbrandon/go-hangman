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

func GetRandomWord(words []string) (string, error) {
	if len(words) == 0 {
		return "", fmt.Errorf("word list is empty")
	}
	return words[rand.IntN(len(words))], nil
}

func CheckValidAlphabet(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func findLetterPosition(letter rune, word string) []int {
	var indexes []int
	for idx, char := range word {
		if letter == char {
			indexes = append(indexes, idx)
			// fmt.Printf("found at index %d, letter %c\n", idx, char)
		}
	}
	return indexes
}

// maskLetters  is used to mask the word with '_' except the first and last letter, and return the masked word.
func maskLetters(word string) string {
	runes := []rune(word)
	length := len(runes)

	// we may not need this, because a word of 2 letters isn't worth playing.
	if length <= 2 {
		return word
	}

	for i := 1; i < length-1; i++ {
		runes[i] = '_'
	}

	return string(runes)
}

func unmaskLetters(word string, letter rune, positions []int) string {
	runes := []rune(word)

	for _, pos := range positions {
		runes[pos] = letter
	}

	return string(runes)
}

func playGame(wordList []string) {
	for {
		randomWord, err := GetRandomWord(wordList)
		if err != nil {
			fmt.Println("Error occurred while fetching a random word:", err)
			return
		}
		const bonusTry = 2
		maxGuessCount := len(randomWord) + bonusTry

		fmt.Printf("%d letters, max try %d, enter a letter or 'quit' to exit.\n\n", len(randomWord), maxGuessCount)
		guessedWord := maskLetters(randomWord)
		guessedLetters := make(map[rune]bool)
		reader := bufio.NewReader(os.Stdin)

		for i := 0; i < maxGuessCount; i++ {
			// print current word
			fmt.Printf("word: %s\n", guessedWord)

			// ask for input
			fmt.Printf("You have %d tries left. guess a letter => ", maxGuessCount-i)

			// read input
			input, _ := reader.ReadString('\n')

			// validate input
			trimmedInput := strings.TrimSpace(input)

			if len(trimmedInput) == 0 {
				fmt.Println("invalid alphabet, please use 'A' - 'Z' or 'a' - 'z'")
				continue
			}

			// compare input
			if trimmedInput == "quit" {
				return
			}

			if !CheckValidAlphabet(rune(trimmedInput[0])) {
				fmt.Println("invalid alphabet, please use 'A' - 'Z' or 'a' - 'z'")
				continue
			}

			guessedLetter := rune(trimmedInput[0])
			if guessedLetters[guessedLetter] {
				fmt.Printf("you have already guessed the letter %c.\n", guessedLetter)
				continue
			}

			// push the new letter into guessed letters slice.
			guessedLetters[guessedLetter] = true

			// find the letter in randomWord, and return the position(s).
			positions := findLetterPosition(rune(trimmedInput[0]), randomWord)
			guessedWord = unmaskLetters(guessedWord, rune(trimmedInput[0]), positions)

			// compare with the final answer
			if guessedWord == randomWord {
				fmt.Printf("Congratulations! You guessed the word: %s\n\n", randomWord)
				break
			}

			fmt.Print("guessed letters => ")
			for c := range guessedLetters {
				fmt.Printf("%c ", c)
			}
			fmt.Println()
			fmt.Println()

			if i == maxGuessCount-1 {
				fmt.Printf("Game over! The word was: %s\n\n", randomWord)
			}
		}
	}
}

func main() {
	file, err := OpenDictionary("wordlist.10000.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wordList := FilterWordsByLength(file, 3, 10)
	playGame(wordList)
	fmt.Println("good bye!")

}
