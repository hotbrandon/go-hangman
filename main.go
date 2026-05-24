package main

import (
	"bufio"
	"fmt"
	"os"
)

func OpenDictionary(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// count the number of words whose length is between minLen and maxLen
func GetWordCount(file *os.File, minLen, maxLen int) int {
	scanner := bufio.NewScanner(file)
	var lineCount int

	for scanner.Scan() {
		// Scanner.Text() automatically strips the newline character (\n),
		line := scanner.Text()
		if len(line) >= minLen && len(line) <= maxLen {
			lineCount++
			fmt.Println(line)
		}
	}

	return lineCount
}

func main() {
	file, err := OpenDictionary("words_alpha.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineCount := GetWordCount(file, 3, 12)

	fmt.Printf("line count: %d\n", lineCount)
	// guess_count := 6
	// for i := 0; i < guess_count; i++ {
	// 	fmt.Println("Guess the next letter:")
	// 	reader := bufio.NewReader(os.Stdin)

	// 	input, _ := reader.ReadString('\n')
	// 	letter := strings.TrimSpace(input)

	// 	fmt.Println(letter)
	// }

}
