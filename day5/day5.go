package main

import (
	"strconv"
	
	"fmt"
	"bufio"
	"os"
	
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

func stringArrayToIntArray(lines []string) []int {
	var ret []int
	for _, i := range lines {
		conv, _ := strconv.Atoi(i)
		ret = append(ret, conv)
	}
	return ret
}

type jumpmap map[int]int

func runTheMaze(jm jumpmap) int {
	moves := -1
	next  := 0
	lenOfMap := len(jm)
	for {
		// fmt.Println(next, jm)
		if len(jm) > lenOfMap {
			return moves
		}
		nextNext := jm[next] + next
		jm[next] = jm[next] + 1
		next = nextNext
		moves++
	}
}

func runTheMaze2(jm jumpmap) int {
	moves := -1
	next  := 0
	lenOfMap := len(jm)
	for {
		// fmt.Println(next, jm)
		if len(jm) > lenOfMap {
			return moves
		}
		nextNext := jm[next] + next
		if jm[next] >= 3 {
			jm[next] = jm[next] - 1
	 	} else {
			jm[next] = jm[next] + 1
		}
		next = nextNext
		moves++
	}
}

func newJumpMap(ints []int) jumpmap {
	tjm := make(jumpmap)
	for i, j := range ints {
		tjm[i] = j
	}
	return tjm
}

func partOne() {
	testInts := []int{0, 3, 0, 1, -3}
	
		lines := readFile("puzzleInput.txt")
		ints := stringArrayToIntArray(lines)
	
		tjm := newJumpMap(testInts)
		jm  := newJumpMap(ints)
	
		fmt.Println(runTheMaze(tjm))
		fmt.Println(runTheMaze(jm))
}

func partTwo() {
	testInts := []int{0, 3, 0, 1, -3}

		lines := readFile("puzzleInput.txt")
		ints  := stringArrayToIntArray(lines)

		tjm := newJumpMap(testInts)
		jm  := newJumpMap(ints)

		fmt.Println(runTheMaze2(tjm))
		fmt.Println(runTheMaze2(jm))
}

func main() {
	partOne()
	//answer 5, 387096
	partTwo()
	//answer 10, 28040648
}