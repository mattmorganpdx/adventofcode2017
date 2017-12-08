package day6

import (
	"fmt"
)

func mem(input ...int) (int, int) {
	count := 0
	seen  := make(map[string]bool)
	returnCount := 0
	var loop string
	loopCount := -1

	for {
		largest := 0
		for i := len(input) - 1 ; i >= 0 ; i-- {
			if input[i] >= input[largest] {
				largest = i
			}
		}
		redist := input[largest]
		input[largest] = 0
		
		for redist > 0{
			if largest + 1 >= len(input) {
				largest = 0
			} else {
				largest++
			}
			input[largest]++
			redist--
			
		}
		count++
		if loopCount >= 0 {
			loopCount++
		}
		if seen[fmt.Sprint(input)] {
			if loop == fmt.Sprint(input){
				return returnCount, loopCount
			} else if loop == "" {
				loop = fmt.Sprint(input)
				loopCount = 0
				returnCount = count
			}
		} 
		
		seen[fmt.Sprint(input)] = true				
	}
}