package day11

import (
	"fmt"
	"strings"

)

type Hex struct {
	x, y, z int
}

func Add(*h Hex, hex Hex) Hex {
	var newHex Hex
	newHex.x = h.x + hex.x
	newHex.y = h.y + hex.y
	newHex.z = h.z + hex.z
	return newHex
}

type MoveMap map[string]Hex

func day11(input string) int {
	path := strings.Split(input, ",")
	hex := Hex{
		x:  0,
		y:  0,
		z:  0,
	}

	moveMap := MoveMap{
		"n": Hex{0, 1, -1},
		"ne": Hex{1, 1, 0},
	}

	for _, p := range path {
		fmt.Println(moveMap)
		hex := hex.Add(moveMap[p])
		fmt.Println(p, hex)
	}

	return 0
}