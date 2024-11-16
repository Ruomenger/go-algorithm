package sliding_window

import (
	"reflect"
	"testing"
)

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abciiidef", 3}, 3},
		{"test2", args{"aeiou", 2}, 2},
		{"test3", args{"leetcode", 3}, 2},
		{"test4", args{"rhythms", 4}, 0},
		{"test5", args{"weallloveyou", 7}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{"cbaebabacd", "abc"}, []int{0, 6}},
		{"test2", args{"abab", "ab"}, []int{0, 1, 2}},
		{"test3", args{"acdcaeccde", "c"}, []int{1, 3, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
