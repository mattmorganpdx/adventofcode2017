package day9

import (
	"strings"
)

func day9(input string) (int, int) {
	depth := 0
	score := 0
	inGarbage := false
	garbageCount := 0
	chars := strings.Split(input, "")
	/* Process the input one char at a time.
	   Inc and dec depth as groups are are found and closed.
	   Skip over any char after "!" by increasing the index an extra time.
	   Add up all the chars found inside garbage along the way.
	*/
	for i := 0 ; i < len(chars) ; i++ {
		switch chars[i] {
		case "{":
			if !inGarbage {
				depth++
			} else {
				garbageCount++
			}
		case "}":
			if !inGarbage {
				score += depth
				depth--
			} else {
				garbageCount++
			}
		case "<":
			if inGarbage {
				garbageCount++
			} else {
				inGarbage = true
			}
		case ">":
			inGarbage = false
		case "!":
			i++
		default:
			if inGarbage {
				garbageCount++
			}
		}
	}
	return score, garbageCount
}