package day8

import (
	"sort"
	"strconv"
	"fmt"
	"strings"

)

type registers map[string]int 

func day8(input []string) (int, int) {
	regs := make(registers)
	var eval []string
	var act  []string
	max := 0
	for _, line := range input {
		prog := strings.Split(line, " if ")
		act = strings.Split(prog[0], " ")
		eval = strings.Split(prog[1], " ")
		//fmt.Println("action:", act, "evaluation:", eval)
		if evaluate(eval, regs) {
			reg := act[0]
			val, _ := strconv.Atoi(act[2])
			if act[1] == "inc" {
				regs[reg] += val
				if regs[reg] > max {
					max = regs[reg]
				}
			} else if act[1] == "dec" {
				regs[reg] -= val
				if regs[reg] > max {
					max = regs[reg]
				}
			}
		}
	}
	fmt.Println(max)
	return getLargestVal(regs), max
}

func getLargestVal(regs registers) int {
	var i []int
	for _, v := range regs {
		i = append(i, v)
	}
	sort.Ints(i)
	return i[len(i)-1]
	
}

func evaluate(eval []string, regs registers) bool {
	regVal := regs[eval[0]]
	comVal, _ := strconv.Atoi(eval[2])
	switch eval[1] {
	case ">":
		if regVal > comVal {
			return true
		}
	case "<":
		if regVal < comVal {
			return true
		}
	case ">=":
		if regVal >= comVal {
			return true
		}
	case "<=":
		if regVal <= comVal {
			return true
		}
	case "==":
		if regVal == comVal {
			return true
		}
	case "!=":
		if regVal != comVal {
			return true
		}	
	}
	return false
}