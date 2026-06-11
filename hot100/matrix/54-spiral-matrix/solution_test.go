package spiral_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	type testCase struct {
		name     string
		matrix   [][]int
		expected []int
	}

	tests := []testCase{
		{
			name: "Example 1 - 3x3",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "Example 2 - 3x4",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "Single Row",
			matrix: [][]int{
				{1, 2, 3},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Single Column",
			matrix: [][]int{
				{1},
				{2},
				{3},
			},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty Matrix",
			matrix:   [][]int{},
			expected: []int{},
		},
		{
			name: "Single Element",
			matrix: [][]int{
				{5},
			},
			expected: []int{5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := spiralOrder(tc.matrix)
			assert.Equal(t, tc.expected, result)
		})
	}
}
