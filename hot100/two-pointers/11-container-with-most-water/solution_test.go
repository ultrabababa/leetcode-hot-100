package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	type testCase struct {
		name     string
		height   []int
		expected int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49, // 选取 8 和 7 (下标 1 和 8)，宽 7，高 7 => 49
		},
		{
			name:     "Example 2 - Two Lines",
			height:   []int{1, 1},
			expected: 1, // 宽 1，高 1 => 1
		},
		{
			name:     "U Shaped",
			height:   []int{4, 3, 2, 1, 4},
			expected: 16, // 头尾两个 4，宽 4，高 4 => 16
		},
		{
			name:     "Increasing",
			height:   []int{1, 2, 3, 4, 5, 25, 24},
			expected: 24, // 最后两个柱子，宽 1，高 24 => 24
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maxArea(tc.height)
			assert.Equal(t, tc.expected, result)
		})
	}
}
