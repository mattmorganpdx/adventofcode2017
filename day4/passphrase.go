package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

// Try and read all lines from puzzleInput.txt and don't care about any errors
func readFile() []string {
	file, _ := os.Open("puzzleInput.txt")

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Determine if passphrase has any duplicate words
func isValid(passphrase string) bool {
	wordHash := make(map[string]bool)
	words := strings.Split(passphrase, " ")
	for _, i := range words {
		if !wordHash[i] {
			wordHash[i] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	lines := readFile()

	// To track the return
	validCount := 0

	for _, line := range lines {
		if isValid(line) {
			validCount++
		}
	}

	fmt.Println(validCount)
}