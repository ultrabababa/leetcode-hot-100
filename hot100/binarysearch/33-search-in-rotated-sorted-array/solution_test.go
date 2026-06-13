package main

import (
	"testing"
)

func TestSearchRotatedArray(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1: Target exists in the left sorted half",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "Example 2: Target does not exist",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "Example 3: Single element, target not found",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			name:   "Edge Case: Single element, target found",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "Edge Case: Two elements, target at index 1",
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
		{
			name:   "Edge Case: Array not actually rotated (k=0)",
			nums:   []int{1, 2, 3, 4, 5},
			target: 4,
			want:   3,
		},
		{
			name:   "Edge Case: Target is the rotation boundary",
			nums:   []int{5, 1, 3},
			target: 5,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
