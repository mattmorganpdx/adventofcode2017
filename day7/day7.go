package day7

import (
	"strings"
	"os"
	"bufio"
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

func getBase() string {
	lines := readFile()
	kidsMap := make(map[string]bool)
	var bases []string
	for _, line := range lines {
		if strings.Contains(line, "->") {
			subs := strings.Split(line, "->")
			bases = append(bases, strings.TrimSpace(strings.Split(subs[0], " ")[0]))
			for _, kid := range strings.Split(subs[1], ",") {
				kidsMap[strings.TrimSpace(kid)] = true
			}
		}
	}

	for _, base := range bases {
		if !kidsMap[base] {
			return base
		}
	}
	return ""
}