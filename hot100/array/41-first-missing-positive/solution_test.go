package first_missing_positive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		expected int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 0},
			expected: 3, // [1, 2] 都有了，缺 3
		},
		{
			name:     "Example 2",
			nums:     []int{3, 4, -1, 1},
			expected: 2, // 有 1，缺 2
		},
		{
			name:     "Example 3",
			nums:     []int{7, 8, 9, 11, 12},
			expected: 1, // 缺 1
		},
		{
			name:     "Sorted 1 to N",
			nums:     []int{1, 2, 3},
			expected: 4, // 1-3 都有，缺 4
		},
		{
			name:     "Single Element 1",
			nums:     []int{1},
			expected: 2,
		},
		{
			name:     "Single Element 2",
			nums:     []int{2},
			expected: 1,
		},
		{
			name:     "With Duplicates",
			nums:     []int{1, 1},
			expected: 2,
		},
		{
			name:     "All Negatives",
			nums:     []int{-5, -2, -1},
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 注意：这题允许修改原数组，所以不需要拷贝
			result := firstMissingPositive(tc.nums)
			assert.Equal(t, tc.expected, result)
		})
	}
}
