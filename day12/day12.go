package day12

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"bufio"
)

// Try and read all lines from puzzleInput.txt and don't care about any errors
func readFile(filename string) []string {
	file, _ := os.Open(filename)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func inputToMap(input []string) map[int][]int {
	returnMap := make(map[int][]int)
	for _, line := range input {
		parts := strings.Split(line, "<->")
		
		var i []int
		for _, e := range strings.Split(parts[1], ",") {
			convertedInt, _ := strconv.Atoi(strings.Trim(e, " "))
			
			i = append(i,convertedInt)
		}
		convertedKey, _ := strconv.Atoi(strings.Trim(parts[0]," "))
		returnMap[convertedKey] =  i
	}
	return returnMap
}

func getConnections(id int, inputMap map[int][]int, returnMap map[int]bool) map[int]bool {
	
	for _, i := range inputMap[id] {
		if returnMap[i] == false {
			returnMap[i] = true
			getConnections(i , inputMap, returnMap)
		}
	}
	return returnMap
}

func day12(id int, inputMap map[int][]int) int {
	_ = fmt.Sprintf("format string")
	r := getConnections(id, inputMap, make(map[int]bool))
	return len(r)
}