package day7

import (
	"strconv"
	"strings"
	"os"
	"bufio"
	"fmt"
	"regexp"
)

// Try and read all lines from puzzleInput.txt and don't care about any errors
func readFile() []string {
	file, _ := os.Open("puzzleInput.txt")

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Parent struct to hold data on base nodes
type Parent struct {
	node string
	childern []string
	childWeights int
}

func getBase() string {
	lines := readFile()
	// map of weights
	towerMap := getWeightMap(lines)
	// Do you have a parent 
	kidsMap := make(map[string]bool)
	// the total weight of all the childern
	kidWeightMap := make(map[string]int)
	
	// the slice of parents
	bases := make(map[string]Parent)
	for _, line := range lines {
		if strings.Contains(line, "->") {
			// var weight int
			subs := strings.Split(line, "->")
			trimBase := strings.TrimSpace(strings.Split(subs[0], " ")[0])
			//weight +=towerMap[trimBase]
			parent := Parent{
				trimBase,
				[]string{},
				0,
			}
			for _, kid := range strings.Split(subs[1], ",") {
				trimKid := strings.TrimSpace(kid)
				parent.childern = append(parent.childern, trimKid)
				parent.childWeights += towerMap[trimKid]
				kidsMap[trimKid] = true
			}
			kidWeightMap[trimBase] = parent.childWeights
			bases[parent.node] = parent
		}
	}


	

	// The Bottom is the one parent that hasn't been seen as a child
	var theBottom string	
	for _, base := range bases {
		if !kidsMap[base.node] {
			theBottom = base.node
		}
		var tw []int
		for _, k := range base.childern {
			tw = append(tw, totalWeight(k, bases, towerMap))
		}
		for _, i := range tw {
			if tw[0] != i {
				fmt.Println("uneveness detected: ", base)
				fmt.Println(base.childern, tw)
				for _, j := range base.childern {
					fmt.Println(j, towerMap[j])
				}
			}
		}

	}					 
	// PART solution node cwwwj needs to be 201 - 8 = 193
	fmt.Println(theBottom)
	return theBottom
}

func totalWeight(node string, bases map[string]Parent, towerMap map[string]int) int {
	var ret int
	if b, ok := bases[node] ; ok {
		for _, kid := range b.childern {
			ret += totalWeight(kid, bases, towerMap)
		}
	}
	ret += towerMap[node]
	return ret
}

func getWeightMap(tower []string) map[string]int {
	towerMap := make(map[string]int)
	for _, level := range tower {
		r := regexp.MustCompile(`(?P<name>\D*) \((?P<weight>\d*)`)
		results := r.FindStringSubmatch(level)
		val, _ := strconv.Atoi(results[2])
		towerMap[strings.TrimSpace(results[1])] = val
	}
	return towerMap
}