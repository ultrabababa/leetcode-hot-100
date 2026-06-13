package main

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "Example 1: Target exists",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			name:   "Example 2: Target does not exist",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		{
			name:   "Edge Case: 1x1 matrix, target found",
			matrix: [][]int{{1}},
			target: 1,
			want:   true,
		},
		{
			name:   "Edge Case: 1x1 matrix, target not found",
			matrix: [][]int{{1}},
			target: 2,
			want:   false,
		},
		{
			name:   "Edge Case: Target smaller than all elements",
			matrix: [][]int{{1, 3, 5}, {7, 9, 11}},
			target: 0,
			want:   false,
		},
		{
			name:   "Edge Case: Target larger than all elements",
			matrix: [][]int{{1, 3, 5}, {7, 9, 11}},
			target: 20,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
