package rotate_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		k        int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			name:     "K is larger than length",
			nums:     []int{1, 2},
			k:        3, // 相当于 k=1: [2, 1]
			expected: []int{2, 1},
		},
		{
			name:     "K is zero",
			nums:     []int{1, 2, 3},
			k:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "K equals length",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Single Element",
			nums:     []int{1},
			k:        100,
			expected: []int{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 因为是原地修改，不需要接收返回值
			rotate(tc.nums, tc.k)
			assert.Equal(t, tc.expected, tc.nums)
		})
	}
}
