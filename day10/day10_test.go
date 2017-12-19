package day10

import "testing"

func Test_day10(t *testing.T) {
	type args struct {
		size  int
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{256, "1,2,3"},
			"3efbe78a8d82f29979031a4aa0b16a9d",
		},
		{
			"test2",
			args{256, "AoC 2017"},
			"33efeb34ea91902bb2f59c9920caa6cd",
		},
		{
			"test3",
			args{256, "1,2,4"},
			"63960835bcdc130f0b66d7ff4f6a5a8e",
		},
		{
			"test4",
			args{256, "106,118,236,1,130,0,235,254,59,205,2,87,129,25,255,118"},
			"9d5f4561367d379cfbf04f8c471c0095",
		},
		{
			"test5",
			args{256, ""},
			"a2582a3a0e66e6e86e3812dcb672a272",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day10(tt.args.size, tt.args.input); got != tt.want {
				t.Errorf("day10() = %v, want %v", got, tt.want)
			}
		})
	}
}
