package array

import (
	"reflect"
	"testing"
)

// https://github.com/youngyangyang04/leetcode-master?tab=readme-ov-file
func TestRemoveElement(t *testing.T) {
	cases := []struct {
		description string
		nums        []int
		val         int
		answers     []int
		length      int
	}{
		{
			description: "example0",
			nums:        []int{1, 2, 3, 3},
			val:         0,
			answers:     []int{1, 2, 3, 3},
			length:      4,
		},
		{
			description: "example1",
			nums:        []int{3, 2, 2, 3},
			val:         3,
			answers:     []int{2, 2},
			length:      2,
		},
		{
			description: "example2",
			nums:        []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:         2,
			answers:     []int{0, 1, 3, 0, 4},
			length:      5,
		},
		{
			description: "example3",
			nums:        []int{0},
			val:         -1,
			answers:     []int{0},
			length:      1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := removeElement(tt.nums, tt.val)
			if result != tt.length {
				t.Errorf("expected length is %d, got %d!", tt.length, result)
			}
		})
	}
}

func TestMid(t *testing.T) {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	nums := []int{-5, 0, 1, 1, 2}
	left := 0
	right := len(nums) - 1
	for left < right-1 {
		mid := left + (right-left)>>1
		if abs(nums[left]) < abs(nums[right]) {
			if nums[right]*nums[mid] > 0 {
				right = mid
			} else {
				left = mid
			}
		} else {
			if nums[left]*nums[mid] > 0 {
				left = mid
			} else {
				right = mid
			}
		}
	}
	t.Logf("left=%d,right=%d", left, right)
}

func TestSortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example0",
			args: args{[]int{-4, -1, 0, 3, 10}},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "example1",
			args: args{[]int{-7, -3, 2, 3, 11}},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "example2",
			args: args{[]int{0}},
			want: []int{0},
		},
		{
			name: "example3",
			args: args{[]int{-5, 0, 1, 1, 2}},
			want: []int{0, 1, 1, 4, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example0",
			args: args{7, []int{2, 3, 1, 2, 4, 3}},
			want: 2,
		},
		{
			name: "example1",
			args: args{4, []int{1, 4, 4}},
			want: 1,
		},
		{
			name: "example2",
			args: args{11, []int{1, 1, 1, 1, 1, 1, 1, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example0",
			args: args{1},
			want: [][]int{{1}},
		},
		{
			name: "example1",
			args: args{3},
			want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
