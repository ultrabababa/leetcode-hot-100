package number_of_islands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	// 辅助函数：将字符串数组转为 byte 的二维切片
	toGrid := func(strs []string) [][]byte {
		grid := make([][]byte, len(strs))
		for i, s := range strs {
			grid[i] = []byte(s)
		}
		return grid
	}

	type testCase struct {
		name     string
		input    []string
		expected int
	}

	tests := []testCase{
		{
			name: "Example 1: Single big island",
			input: []string{
				"11110",
				"11010",
				"11000",
				"00000",
			},
			expected: 1,
		},
		{
			name: "Example 2: Multiple islands",
			input: []string{
				"11000",
				"11000",
				"00100",
				"00011",
			},
			expected: 3,
		},
		{
			name:     "Empty grid",
			input:    []string{},
			expected: 0,
		},
		{
			name: "All water",
			input: []string{
				"000",
				"000",
			},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			grid := toGrid(tc.input)
			assert.Equal(t, tc.expected, numIslands(grid))
		})
	}
}
