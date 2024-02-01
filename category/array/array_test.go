package array

import (
	"testing"
)

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
