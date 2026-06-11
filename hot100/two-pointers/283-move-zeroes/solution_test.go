package move_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Example 2 - Single Zero",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "Single Non-Zero",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "No Zeros",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "All Zeros",
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "Zeros at End",
			nums:     []int{4, 2, 0, 0},
			expected: []int{4, 2, 0, 0},
		},
		{
			name:     "Zeros at Start",
			nums:     []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 注意：这里不需要接收返回值，因为是原地修改
			moveZeroes(tc.nums)

			// 直接检查 tc.nums 是否变成了我们要的样子
			assert.Equal(t, tc.expected, tc.nums, "Slice should be modified in-place")
		})
	}
}
