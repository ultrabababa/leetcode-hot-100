package main

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1: Target exists multiple times",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "Example 2: Target does not exist",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "Example 3: Empty array",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			name:   "Edge Case: Target at the very beginning",
			nums:   []int{2, 2, 3, 4, 5},
			target: 2,
			want:   []int{0, 1},
		},
		{
			name:   "Edge Case: Target at the very end",
			nums:   []int{1, 2, 3, 5, 5},
			target: 5,
			want:   []int{3, 4},
		},
		{
			name:   "Edge Case: Single element, target found",
			nums:   []int{8},
			target: 8,
			want:   []int{0, 0},
		},
		{
			name:   "Edge Case: All elements are target",
			nums:   []int{8, 8, 8, 8},
			target: 8,
			want:   []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
