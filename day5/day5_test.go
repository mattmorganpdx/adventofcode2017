package day5

import (
	"testing"
	"fmt"
)

func Test_runTheMaze(t *testing.T) {
	type args struct {
		jm jumpmap
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{	"test1",
			args{newJumpMap([]int{0, 3, 0, 1, -3})},
			5,
		},
		{	"test2",
			args{newJumpMap(stringArrayToIntArray(readFile("puzzleInput.txt")))},
			387096,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runTheMaze(tt.args.jm); got != tt.want {
				t.Errorf("runTheMaze() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runTheMaze2(t *testing.T) {
	type args struct {
		jm jumpmap
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{	"test1",
			args{newJumpMap([]int{0, 3, 0, 1, -3})},
			10,
		},
		{	"test2",
			args{newJumpMap(stringArrayToIntArray(readFile("puzzleInput.txt")))},
			28040648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runTheMaze2(tt.args.jm); got != tt.want {
				t.Errorf("runTheMaze2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Example_runTheMaze() {
	jm := newJumpMap([]int{0, 3, 0, 1, -3})

	fmt.Println(runTheMaze(jm))
	// Output: 5
}