package sliding_window

import "testing"

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
