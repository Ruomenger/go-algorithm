package hashtable

import (
	"reflect"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example0", args{"abc", "bca"}, true},
		{"example1", args{"rat", "car"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonChars(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"example1", args{[]string{"bella", "label", "roller"}}, []string{"e", "l", "l"}},
		{"example2", args{[]string{"cool", "lock", "cook"}}, []string{"c", "o"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2}},
		{"example2", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{9, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{19}, true},
		{"example1", args{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"example1", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"example1", args{[]int{3, 3}, 6}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourSumCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}}, 2},
		{"example2", args{[]int{0}, []int{0}, []int{0}, []int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.nums1, tt.args.nums2, tt.args.nums3, tt.args.nums4); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
