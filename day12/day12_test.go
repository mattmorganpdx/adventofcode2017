package day12

import "testing"

func Test_day12(t *testing.T) {
	testInput := []string{
		"0 <-> 2",
		"1 <-> 1",
		"2 <-> 0, 3, 4",
		"3 <-> 2, 4",
		"4 <-> 2, 3, 6",
		"5 <-> 6",
		"6 <-> 4, 5",
	}
	type args struct {
		id int
		inputMap map[int][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{	"test1",
			args{0, inputToMap(testInput)},
			6,
		},
		{	"test2",
			args{0, inputToMap(readFile("puzzleInput.txt"))},
			175,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day12(tt.args.id, tt.args.inputMap); got != tt.want {
				t.Errorf("day12() = %v, want %v", got, tt.want)
			}
		})
	}
}
