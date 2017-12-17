package day10

import "testing"

func Test_day10(t *testing.T) {
	type args struct {
		size  int
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{5, []int{3, 4, 1, 5}},
			12,
		},
		{
			"test2",
			args{256, []int{106,118,236,1,130,0,235,254,59,205,2,87,129,25,255,118}},
			6909,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day10(tt.args.size, tt.args.input...); got != tt.want {
				t.Errorf("day10() = %v, want %v", got, tt.want)
			}
		})
	}
}
