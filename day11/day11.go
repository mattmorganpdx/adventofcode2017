package day11

import (
	"fmt"
	"strings"

)

func day11(input string) int {
	path := strings.Split(input, ",")

	for _, p := range path {
		fmt.Println(p)
	}

	return 0
}