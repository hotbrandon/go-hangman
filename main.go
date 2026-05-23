package main

import (
	"bufio"
	"fmt"
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

func main() {
	file, err := OpenDictionary("words_alpha.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	guess_count := 6
	for i := 0; i < guess_count; i++ {
		fmt.Println("Guess the next letter:")
		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')
		letter := strings.TrimSpace(input)

		fmt.Println(letter)
	}

}
