package day6

import "testing"

func Test_mem(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want1 int
		want2 int 
	}{
		{	"test1",
			args{[]int{0,2,7,0}},
			5, 4,
		},
		{	"test2",
			args{[]int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}},
			7864, 1695,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, get := mem(tt.args.input...); got != tt.want1 || get != tt.want2 {
				t.Errorf("mem() = %v, %v, want1 %v, want2 %v ", got, get, tt.want1, tt.want2)
			}
		})
	}
}
