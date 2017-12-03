package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type Table [][]int

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

func textToTable(text []string) Table {
	var ret Table
	for _, line := range text {
		var row []int
		for _, elem := range strings.Split(line,"\t") {
			intElem, _ := strconv.Atoi(elem)
			row = append(row, intElem)
		}
		ret = append(ret, row)
	}

	return ret
}

func main () {
	lines := readFile()

	checkTable := textToTable(lines)
	
}