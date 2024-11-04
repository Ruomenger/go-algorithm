package greedy

import (
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3}, []int{1, 1}}, 1},
		{"test2", args{[]int{1, 2}, []int{1, 2, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWiggleMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 7, 4, 9, 2, 5}}, 6},
		{"test2", args{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}}, 7},
		{"test3", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2},
		{"test3", args{[]int{0, 0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
