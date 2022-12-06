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

func Test_digitize(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"digitize", args{9078562341}, []int{9, 0, 7, 8, 5, 6, 2, 3, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitize(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digitze() = %v, want %v", got, tt.want)
			}
			if got := appendDigitze(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendDigitze() = %v, want %v", got, tt.want)
			}
			if got := copyDigitize(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyDigitze() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_digitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		digitize(9078562341)
	}
}
func Benchmark_appendDigitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendDigitze(9078562341)
	}
}
func Benchmark_copyDigitize(b *testing.B) { // winner winner tofu dinner.
	for i := 0; i < b.N; i++ {
		copyDigitize(9078562341)
	}
}
