package findall

import (
	"reflect"
	"testing"
)

func Test_findAll(t *testing.T) {
	type args struct {
		sumOfDigits    int
		numberOfDigits int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{10, 3}, []int{8, 118, 334}},
		{"case2", args{27, 3}, []int{1, 999, 999}},
		{"doesnt satsify constraints, should return nil", args{84, 4}, []int{}},
		{"case3", args{35, 6}, []int{123, 116999, 566666}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAll(tt.args.sumOfDigits, tt.args.numberOfDigits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
