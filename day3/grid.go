package main

import (
	"fmt"
	"math"
)

/*
17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...


1: 0,0
2: 1,0 R1
3: 1,1 U1
4: 0,1 L1
5: -1,1 L2
6: -1,0 D1
7: -1,-1 D2
8: 0, -1 R1
9: 1, -1 R2
10: 2, -1 R3

 */

// Direction object to control moves
type Direction struct {
	max int
	counter int
	nextDirection *Direction
	move func(square Coordinate) Coordinate
}

// Good ol' x, y
type Coordinate struct {
	x int
	y int
}

// The grid we're creating
type Grid map[int]Coordinate

// Numbers based on Coordinates
type Dirg map[Coordinate]int

// These 4 functions return the coordinate of the square in that direction
func Right(this Coordinate) Coordinate {
	var newSquare Coordinate
	newSquare.x = this.x + 1
	newSquare.y = this.y
	return newSquare
}

func Left(this Coordinate) Coordinate {
	var newSquare Coordinate
	newSquare.x = this.x - 1
	newSquare.y = this.y
	return newSquare
}

func Up(this Coordinate) Coordinate {
	var newSquare Coordinate
	newSquare.x = this.x
	newSquare.y = this.y + 1
	return newSquare
}

func Down(this Coordinate) Coordinate {
	var newSquare Coordinate
	newSquare.x = this.x
	newSquare.y = this.y - 1
	return newSquare
}

func grid1() {
	// Make the grid and define the start 0,0 = #1
	grid := make(Grid)
	grid[1] = Coordinate{0, 0}

	// define all the directions linked to allow spiralling out
	right := Direction{1,1,nil, Right }
	up := Direction{1,1,nil,Up}
	left := Direction{2,2,nil,Left}
	down := Direction{2, 2, nil,Down}

	right.nextDirection = &up
	up.nextDirection = &left
	left.nextDirection = &down
	down.nextDirection = &right

	currentDirection := &right

	// loop from 2 to some value higher than the target
	for i := 2 ; i < 350000 ; i++ {
		grid[i] = currentDirection.move(grid[i - 1])
		currentDirection.counter--
		if currentDirection.counter <= 0 {
			currentDirection.max = currentDirection.max + 2
			currentDirection.counter = currentDirection.max
			currentDirection = currentDirection.nextDirection
		}


	}

	target := grid[347991]

	// Print the abs of the target's x + y which is the vaule we're looking for
	fmt.Println(math.Abs(float64(target.x)) + math.Abs(float64(target.y)))
}

func getEight(c Coordinate, d Dirg) int {
	A := d[Coordinate{c.x +1 , c.y -1}]
	B := d[Coordinate{c.x    , c.y -1}]
	C := d[Coordinate{c.x -1 , c.y -1}]
	D := d[Coordinate{c.x -1 , c.y   }]
	E := d[Coordinate{c.x -1 , c.y +1}]
	F := d[Coordinate{c.x    , c.y +1}]
	G := d[Coordinate{c.x +1 , c.y +1}]
	H := d[Coordinate{c.x +1 , c.y   }]

	// fmt.Println(c,d,A, B, C, D, E, F, G, H, A + B + C + D + E + F + G + H)

	return A + B + C + D + E + F + G + H
}

func grid2() {
	dirg := make(Dirg)
	dirg[Coordinate{0, 0}] = 1

	/*
		3,3 4,3  5,3  x-1,y+1 x,y+1 x+1,y+1
		3,2 4,2  5,2  x-1,y   x,y  x+1, y
		3,1 4,1  5,1  x-1,y-1 x,y-1 x+1, y-1

	
	*/

	grid := make(Grid)
	grid[1] = Coordinate{0, 0}

	// define all the directions linked to allow spiralling out
	right := Direction{1,1,nil, Right }
	up := Direction{1,1,nil,Up}
	left := Direction{2,2,nil,Left}
	down := Direction{2, 2, nil,Down}

	right.nextDirection = &up
	up.nextDirection = &left
	left.nextDirection = &down
	down.nextDirection = &right

	currentDirection := &right

	

	// loop from 2 to some value higher than the target
	for i := 2 ; i < 350000 ; i++ {
		grid[i] = currentDirection.move(grid[i - 1])
		eightVal := getEight(grid[i], dirg)
		//fmt.Println(grid[i], dirg, eightVal)
		if eightVal > 347991 {
			fmt.Println(eightVal)
			return
		}
		dirg[grid[i]] = eightVal
		currentDirection.counter--
		if currentDirection.counter <= 0 {
			currentDirection.max = currentDirection.max + 2
			currentDirection.counter = currentDirection.max
			currentDirection = currentDirection.nextDirection
		}


	}

	target := grid[347991]

	// Print the abs of the target's x + y which is the vaule we're looking for
	fmt.Println(math.Abs(float64(target.x)) + math.Abs(float64(target.y)))
}

func main() {
	//480
	grid1()
	//349975
	grid2()
}