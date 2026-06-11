package twosum

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		target   int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1}, // 2 + 7 = 9
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2}, // 2 + 4 = 6
		},
		{
			name:     "Example 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1}, // 3 + 3 = 6
		},
		{
			name:     "Negative Numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4}, // -3 + (-5) = -8
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.nums, tc.target)

			// 断言结果长度必须为 2
			assert.Len(t, result, 2, "Should return exactly 2 indices")

			// 为了保证测试的鲁棒性（不管算法返回 [0,1] 还是 [1,0] 都是对的），
			// 我们先把结果排序，然后再跟期望值对比。
			sort.Ints(result)
			sort.Ints(tc.expected)

			assert.Equal(t, tc.expected, result, "Indices should match")
		})
	}
}
