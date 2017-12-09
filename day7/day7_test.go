package day7

import "testing"

func Test_getBase(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{	"test1",
			"tknk",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBase(); got != tt.want {
				t.Errorf("getBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
