package p0101

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	cases := []struct {
		description string
		str         string
		expected    bool
	}{
		{
			description: "leetcode",
			str:         "leetcode",
			expected:    false,
		},
		{
			description: "abc",
			str:         "abc",
			expected:    true,
		},
		{
			description: "aa",
			str:         "aa",
			expected:    false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := isUniqueBit(tt.str)
			if result != tt.expected {
				t.Errorf("expected %t, but got %t", tt.expected, result)
			}
			result = isUnique(tt.str)
			if result != tt.expected {
				t.Errorf("expected %t, but got %t", tt.expected, result)
			}
		})
	}
}
