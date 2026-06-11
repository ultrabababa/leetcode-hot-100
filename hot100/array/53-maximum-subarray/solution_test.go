package maximum_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		expected int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6, // [4,-1,2,1]
		},
		{
			name:     "Example 2 - Single Element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Example 3 - All Positive",
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23, // 5 + 4 - 1 + 7 + 8
		},
		{
			name:     "All Negative",
			nums:     []int{-2, -5, -9, -1},
			expected: -1, // 最大的是 -1，不是 0
		},
		{
			name:     "Mixed with Zero",
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			name:     "Large Negative Start",
			nums:     []int{-100, 1, 2},
			expected: 3, // 抛弃 -100，从 1 开始
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maxSubArray(tc.nums)
			assert.Equal(t, tc.expected, result)
		})
	}
}
