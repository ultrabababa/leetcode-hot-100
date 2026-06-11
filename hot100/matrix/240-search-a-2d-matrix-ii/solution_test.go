package search_a_2d_matrix_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	type testCase struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}

	tests := []testCase{
		{
			name: "Example 1 - Target exists",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Example 2 - Target does not exist",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target:   20,
			expected: false, // 20 不在矩阵中
		},
		{
			name: "Target smaller than minimum",
			matrix: [][]int{
				{1, 4},
				{2, 5},
			},
			target:   0,
			expected: false,
		},
		{
			name: "Target larger than maximum",
			matrix: [][]int{
				{1, 4},
				{2, 5},
			},
			target:   10,
			expected: false,
		},
		{
			name: "Single Element - Match",
			matrix: [][]int{
				{5},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Single Element - No Match",
			matrix: [][]int{
				{5},
			},
			target:   1,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := searchMatrix(tc.matrix, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}
