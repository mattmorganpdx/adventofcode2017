package day10

import (
	"bytes"
	"fmt"
)

func createArray(size int) []int {
	ret := []int{}
	for i := 0 ; i < size ; i++ {
		ret = append(ret, i)
	}
	return ret
}

func day10(size int, input1 string) string {
	suffix := []int{17, 31, 73, 47, 23}
	var input []int
	for _, r := range []rune(input1){
		input = append(input, int(r)	)
	}
	input = append(input, suffix...)

	position := 0
	skip     := 0
	intRange := createArray(size)

	for i := 0 ; i < 64 ; i++ {
		intRange, position, skip = knot(input, intRange, position, skip)		
	}

	var denseHash []int
	for i := 0 ; i < 256 ; i += 16 {
		bytexor := intRange[i]
		for j := 1; j < 16 ; j++ {
			bytexor = bytexor ^ intRange[i+j]
		}
		denseHash = append(denseHash, bytexor)
	}

	var denseHex bytes.Buffer
	for _, i := range denseHash {
		denseHex.WriteString(fmt.Sprintf("%02x",i))
	}

	return denseHex.String()
}

func knot(input, intRange []int, position, skip int) ([]int, int, int) {
	for i := 0 ; i < len(input) ; i++ {
		length     := input[i]
		rangeStart := position
		rangeEnd   := rangeStart + length - 1
		if rangeEnd < len(intRange) {
			intRange = switcher(rangeStart, rangeEnd, intRange)

		} else {
			front := rangeEnd - len(intRange)
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
		for position > len(intRange) {
			position -= len(intRange)
		}
		skip++
	}
	return intRange, position, skip
}

func switcher(start, end int, intRange []int) []int{
	for f, b  := start, end; f < b ; f, b = f+1, b-1 {
		intRange[f], intRange[b] = intRange[b], intRange[f]
	}
	return intRange
}