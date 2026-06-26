package main

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1: Rotated 3 times",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "Example 2: Rotated 4 times",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "Example 3: Not rotated (or rotated n times)",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "Edge Case: Single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Edge Case: Two elements, rotated",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "Edge Case: Two elements, not rotated",
			nums: []int{1, 2},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
