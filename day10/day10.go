package day10

import (
	
	"fmt"

)

func createArray(size int) []int {
	ret := []int{}
	for i := 0 ; i < size ; i++ {
		ret = append(ret, i)
	}
	return ret
}

func day10(size int, input ...int) int {	
	//fmt.Println(input)

	position := 0
	skip     := 0
	intRange := createArray(size)

	for i := 0 ; i < len(input) ; i++ {
		length     := input[i]
		rangeStart := position
		rangeEnd   := rangeStart + length - 1
		//fmt.Println("start:",rangeStart," end:",rangeEnd)
		if rangeEnd < len(intRange) {
			intRange = switcher(rangeStart, rangeEnd, intRange)

		} else {
			// [0 1 2 3 4]  3 + 4 = 6, (6 - 1 ) - 5
			front := rangeEnd - len(intRange)
			//fmt.Println("front:",front)
			sub := []int{}
			for sr := rangeStart ; sr < len(intRange) ; sr++ {
				sub = append(sub, intRange[sr])
			}
			for sr := 0; sr <= front ; sr++ {
				sub = append(sub, intRange[sr])
			}
			sub = switcher(0, len(sub) - 1, sub)
			for sr := rangeStart ; sr < len(intRange) ; sr++ {
				intRange[sr], sub = sub[0], sub[1:]
			}
			for sr := 0; sr <= front ; sr++ {
				intRange[sr], sub = sub[0], sub[1:]
			}
		}
		position += length + skip
		if len(intRange) < position{
			position -= len(intRange)
		}
		skip++
	}
	fmt.Println(intRange)
	return intRange[0] * intRange[1]
}

func switcher(start, end int, intRange []int) []int{
	//fmt.Println("switcher", start, end, intRange)
	for f, b  := start, end; f < b ; f, b = f+1, b-1 {
		intRange[f], intRange[b] = intRange[b], intRange[f]
	}
	//fmt.Println("switcher", intRange)
	return intRange
}