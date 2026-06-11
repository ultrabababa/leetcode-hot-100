package set_matrix_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	type testCase struct {
		name     string
		matrix   [][]int
		expected [][]int
	}

	tests := []testCase{
		{
			name: "Example 1",
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "Example 2 - Zeros on edges",
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			name: "First Row Zero",
			matrix: [][]int{
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "First Column Zero",
			matrix: [][]int{
				{0, 1},
				{1, 1},
			},
			expected: [][]int{
				{0, 0},
				{0, 1},
			},
		},
		{
			name: "Origin Zero",
			matrix: [][]int{
				{0, 1},
				{1, 1},
			},
			expected: [][]int{
				{0, 0},
				{0, 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 原地修改，不需要接收返回值
			setZeroes(tc.matrix)
			assert.Equal(t, tc.expected, tc.matrix)
		})
	}
}
