package day11

import (
	"math"
	"sort"
	"strings"
)

// Hex struct 
// see this link for more info on how hex grids map to 3D grids
// https://www.redblobgames.com/grids/hexagons/
type Hex struct {
	x, y, z int
}

// Add a hex objects value to the current objects value
func (h *Hex) Add(hex Hex) Hex {
	var newHex Hex
	newHex.x = h.x + hex.x
	newHex.y = h.y + hex.y
	newHex.z = h.z + hex.z
	return newHex
}

// Max finds the largest dimension of the hex 
func (h *Hex) Max() int {
	s := []int{
		int(math.Abs(float64(h.x))),
		int(math.Abs(float64(h.y))),
		int(math.Abs(float64(h.z))),
	}
	sort.Ints(s)
	return s[len(s) - 1]
}

func day11(input string) (int, int) {
	path := strings.Split(input, ",")

	// Start at 0,0,0
	hex := Hex{
		x:  0,
		y:  0,
		z:  0,
	}

	// Map directions to values
	moveMap := map[string]Hex {
		"n":  Hex{0, 1, -1},
		"ne": Hex{1, 0, -1},
		"se": Hex{1, -1, 0},
		"s":  Hex{0, -1, 1},
		"sw": Hex{-1, 0, 1},
		"nw": Hex{-1, 1, 0},
	}

	// Track the farthest move
	farthestMove := hex.Max()

	for _, p := range path {
		hex = hex.Add(moveMap[p])
		if hex.Max() > farthestMove {
			farthestMove = hex.Max()
		}
	
	}

	return hex.Max(), farthestMove
}