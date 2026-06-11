package longestconsecutive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		expected int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4, // [1, 2, 3, 4]
		},
		{
			name:     "Example 2 - Mixed Order",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9, // [0, 1, 2, 3, 4, 5, 6, 7, 8]
		},
		{
			name:     "Example 3 - With Duplicates",
			nums:     []int{1, 0, 1, 2},
			expected: 3, // [0, 1, 2] -> 重复的 1 不计算在长度内
		},
		{
			name:     "Empty Array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single Element",
			nums:     []int{10},
			expected: 1,
		},
		{
			name:     "Negative Numbers",
			nums:     []int{-5, -2, -3, -4, -10},
			expected: 4, // [-5, -4, -3, -2]
		},
		{
			name:     "No Sequence",
			nums:     []int{1, 3, 5, 7},
			expected: 1, // 任意一个元素本身就是长度为 1 的序列
		},
		{
			name:     "Large Gaps",
			nums:     []int{1, 2, 100, 101, 102},
			expected: 3, // [100, 101, 102]
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := longestConsecutive(tc.nums)
			assert.Equal(t, tc.expected, result, "Should find the correct length of longest consecutive sequence")
		})
	}
}
