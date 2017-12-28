package day11

import (
	"testing"
)

func Test_day11(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{	"test1",
			args{"ne,ne,ne"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day11(tt.args.input)
		})
	}
}
