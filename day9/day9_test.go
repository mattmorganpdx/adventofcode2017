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
		want1 int
		want2 int
	}{
		{
			"test1",
			args{"{{<a!>},{<a!>},{<a!>},{<ab>}}"},
			3, 17,
		},
		{
			"test2",
			args{readFile("puzzleInput.txt")[0]},
			17390, 7825, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := day9(tt.args.input); got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("day9() = %v, want1 %v and %v, want2 %v", got1, tt.want1, got2, tt.want2)
			}
		})
	}
}
