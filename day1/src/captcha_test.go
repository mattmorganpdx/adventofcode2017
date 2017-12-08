package main

import (
	"testing"
	"./intTools"
)

func Test_captcha(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{	"test1",
			args{ []int{1,1,2,2} },
			3,
		},
		{
			"test2",
			args { intTools.StringToIntArray(readFile()[0]) },
			1047,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captcha(tt.args.input); got != tt.want {
				t.Errorf("captcha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_captcha2(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{	"test1",
			args{ []int{1,2,3,4,2,5} },
			4,
		},
		{
			"test2",
			args { intTools.StringToIntArray(readFile()[0]) },
			982,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captcha2(tt.args.input); got != tt.want {
				t.Errorf("captcha2() = %v, want %v", got, tt.want)
			}
		})
	}
}
