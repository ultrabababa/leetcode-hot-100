package rotate_image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	type testCase struct {
		name     string
		matrix   [][]int
		expected [][]int
	}

	tests := []testCase{
		{
			name: "Example 1 - 3x3",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "Example 2 - 4x4",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name: "Single Element",
			matrix: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 原地修改
			rotate(tc.matrix)
			assert.Equal(t, tc.expected, tc.matrix)
		})
	}
}
