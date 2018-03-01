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
		want1 int
		want2 int
	}{
		{	"test1",
			args{0, inputToMap(testInput)},
			6, 2,
		},
		{	"test2",
			args{0, inputToMap(readFile("puzzleInput.txt"))},
			175, 213,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := day12(tt.args.id, tt.args.inputMap); got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("day12() = %v, want %v ; %v, want %v", got1, tt.want1, got2, tt.want2)
			}
		})
	}
}
