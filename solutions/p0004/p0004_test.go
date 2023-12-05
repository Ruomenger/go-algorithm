package p0101

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	cases := []struct {
		description  string
		nums1, nums2 []int
		expected     float64
	}{
		{
			description: "example1",
			nums1:       []int{1, 3},
			nums2:       []int{2},
			expected:    2.0,
		},
		{
			description: "example2",
			nums1:       []int{1, 2},
			nums2:       []int{3, 4},
			expected:    2.5,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := findMedianSortedArrays(tt.nums1, tt.nums2)
			if result != tt.expected {
				t.Errorf("expected %f, but got %f", tt.expected, result)
			}
		})
	}
}
