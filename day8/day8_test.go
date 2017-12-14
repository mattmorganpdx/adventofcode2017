package day8

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

func Test_day8(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want1 int
		want2 int
	}{
		{ "test1",
		  args{readFile("testInput.txt")},
		  1, 10,
		},
		{ "test2",
			args{readFile("puzzleInput.txt")},
			5752, 6366,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := day8(tt.args.input); got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("day8() = %v, want %v ; %v, want %v", got1, tt.want1, got2, tt.want2)
			}
		})
	}
}
