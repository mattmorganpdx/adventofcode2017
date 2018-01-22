package day11

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

func Test_day11(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want1 int
		want2 int
	}{
		{	"test1",
			args{"ne,ne,ne"},
			3, 3,
		},
		{
			"test2",
			args{readFile("puzzleInput.txt")[0]},
			685, 1457,
		},
		{
			"test3",
			args{"se,sw,se,sw,sw"},
			3, 3,
		},
		{
			"test4",
			args{"ne,ne,s,s"},
			2, 2,
		},
		{
			"test5",
			args{"ne,ne,sw,sw"},
			0, 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := day11(tt.args.input); got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("day9() = %v, want1 %v and %v, want2 %v", got1, tt.want1, got2, tt.want2)
			}
		})
	}
}
