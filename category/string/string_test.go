package string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "example1", args: args{[]byte{'h', 'e', 'l', 'l', 'o'}}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{name: "example2", args: args{[]byte{'e', 'l', 'l', 'o'}}, want: []byte{'o', 'l', 'l', 'e'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if reverseString(tt.args.s); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("fourSum() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	str := "str你"
	fmt.Println(len(str))
}

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"abababaababacb", "ababacb"}, 7},
		{"example2", args{"aaaaa", "bba"}, -1},
		{"example3", args{"mississippi", "issip"}, 4},
		{"example4", args{"aabaaabaaac", "aabaaac"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{"abcdefg", 2}, "bacdfeg"},
		{"example1", args{"abcd", 2}, "bacd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{"abab"}, true},
		{"example2", args{"aba"}, false},
		{"example3", args{"abcabcabcabc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
