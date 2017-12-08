package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"sort"
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
		sort.Ints(row)
		ret = append(ret, row)
	}

	return ret
}

func summer(checkTable Table) int {
	sum := 0
	
		for i := 0 ; i < len(checkTable); i++ {
			sum += checkTable[i][len(checkTable[i]) - 1] - checkTable[i][0]
		}
	return sum
}

func div(arr []int) int {
	i := arr
	for idx, j := range i {
		max := len(i)
		for e := 0 ; e < max ; e++ {
			if e != idx {
				right := j % i[e]
				left := i[e] % j
				if right == 0 {
					return j / i[e]
				} else if left == 0 {
					return i[e] / j
				}
			}
		}
	}
	return 0
}

func diver(checkTable Table) int {
	sum := 0
		
		
		for i := 0 ; i < len(checkTable); i++ {
			sum += div(checkTable[i])
		}
	return sum
}

func main () {
	lines := readFile()

	//textToTable returns sorted rows to make the subtraction easier
	checkTable := textToTable(lines)

	sum := summer(checkTable)

	//41887
	fmt.Println(sum)

	div := diver(checkTable)

	//226
	fmt.Println(div)
	
}