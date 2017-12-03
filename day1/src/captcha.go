package main

import (
	"bufio"
	"fmt"
	"os"
	"./intTools"
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

// Do the "captcha" work
func captcha(input []int) int {
	output := 0

	capObj := intTools.NewIntIter(input)

	for i := 0; i < len(input); i++ {
		if capObj.Current() == capObj.Peek() {
			output = output + capObj.Next()
		} else {
			capObj.Next()
		}
	}

	return output
}

func main() {
	input := readFile()
	//convert the first line of puzzleInput into an []int
	theNumbers := intTools.StringToIntArray(input[0])
	ret := captcha(theNumbers)
	fmt.Println(ret)

}
