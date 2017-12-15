package day9

import (
	"testing"
	"os"
	"bufio"
)

// Try and read all lines from puzzleInput.txt and don't care about any errors
func readFile(filename string) []string {
	file, _ := os.Open(filename)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Test_day9(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{"{{<a!>},{<a!>},{<a!>},{<ab>}}"},
			3,
		},
		{
			"test2",
			args{readFile("puzzleInput.txt")[0]},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day9(tt.args.input); got != tt.want {
				t.Errorf("day9() = %v, want %v", got, tt.want)
			}
		})
	}
}
