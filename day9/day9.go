package day9

import (
	"fmt"
	"strings"
	"regexp"
)

func day9(input string) int {
	// count := findStart(strings.Split(input, "") , 0, 0)

	return findGroups(input, 1, 1)

}

func findGroups(input string, count, score int) int {
	count += count + 1
	subs := strings.Split(input[1:len(input)-1], ",")
	// fmt.Println(subs)
	r := regexp.MustCompile(`[{]*(?:|<[^!]*>|<[^>]*[^!]>|[^<>]*)[}]*`)
	for _, sub := range subs {
		for _, m := range r.FindAllString(sub, -1) {
			fmt.Println(sub, count)			
			score += findGroups(m, count , score + 1)
		}
	}
	return score 
}

func findStart(input []string, index, count int) int {
	if index >= len(input) {
		return count
	}
	for index < len(input) {
		if input[index] == "{" {
			return findEnd(input, index + 1, count, false)
		}
		index++
	}
	
	return count
}

func findEnd(input []string, index, count int, garbage bool) int {
	if index >= len(input) {
		return count
	}
	if input[index] == "}" {
		if garbage {
			return count
		} 
		
		return count + 1	
	}
	for index < len(input) {
		switch input[index] {
		case "<":
			if input[index -1] != "!" && !garbage{
				//Start garbage
				garbage = true
				return findEnd(input, index + 1, count , garbage)
			}
		case ">":
			if input[index -1] != "!" && garbage {
				garbage = false
				return findEnd(input, index + 1, count , garbage)
			}		
		case "{":
			// Reset garbage
			return findEnd(input, index + 1, count , false)		
		}
		index++
	}

	return count
}