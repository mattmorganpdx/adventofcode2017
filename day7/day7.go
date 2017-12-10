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

func getBase() string {
	lines := readFile()
	towerMap := getWeightMap(lines)
	kidsMap := make(map[string]bool)
	kidWeightMap := make(map[string]int)
	var bases []string
	for _, line := range lines {
		if strings.Contains(line, "->") {
			var weight int
			subs := strings.Split(line, "->")
			trimBase := strings.TrimSpace(strings.Split(subs[0], " ")[0])
			bases = append(bases, trimBase)
			weight +=towerMap[trimBase]
			for _, kid := range strings.Split(subs[1], ",") {
				trimKid := strings.TrimSpace(kid)
				weight += towerMap[trimKid]
				kidsMap[trimKid] = true
			}
			fmt.Println(trimBase, weight)
			kidWeightMap[trimBase] = weight
		}
	}
	var theBottom string	
	for _, base := range bases {
		if !kidsMap[base] {
			theBottom = base
		} 
		if kidWeightMap[base] > 0 {
			fmt.Println(base, kidWeightMap[theBottom])
		}
	}	
	return theBottom
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