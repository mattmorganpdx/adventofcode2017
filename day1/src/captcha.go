package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"./intIter"
)

type Captcha struct {
	value   []int
	length  int
	current int
}

func NewCaptcha(input []int) Captcha {
	return Captcha{input, len(input), 0}

}

func (this *Captcha) Next() int {
	if this.current+1 >= this.length {
		this.current = 1
	} else {
		this.current++
	}
	return this.value[this.current-1]
}

func (this *Captcha) Peek() int {
	var next int
	if this.current+1 >= this.length {
		next = 0
	} else {
		next = this.current + 1
	}
	return this.value[next]
}

func (this *Captcha) Current() int {
	return this.value[this.current]
}

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

func stringToIntArray(input string) []int {
	chars := []rune(input)
	var ret []int

	for _, i := range chars {
		j, _ := strconv.Atoi(string(i))
		ret = append(ret, j)
	}
	return ret
}

func captcha(input []int) int {
	output := 0

	capObj := intIter.NewIntIter(input)

	for i := 0; i < len(input); i++ {
		if capObj.Current() == capObj.Peek() {
			output = output + capObj.Next()
		} else {
			capObj.Next()
		}
	}

	return output
}

func main() {
	input := readFile()
	theNumbers := stringToIntArray(input[0])
	ret := captcha(theNumbers)
	fmt.Println(ret)

}
