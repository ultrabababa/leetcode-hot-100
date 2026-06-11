package subarray_sum_equals_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		k        int
		expected int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2, // [1,1] (index 0-1), [1,1] (index 1-2)
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2, // [1,2], [3]
		},
		{
			name:     "Negative Numbers",
			nums:     []int{1, -1, 0},
			k:        0,
			expected: 3,
			// 解释:
			// 1. [1, -1] = 0
			// 2. [0] = 0
			// 3. [1, -1, 0] = 0
		},
		{
			name:     "Single Element Match",
			nums:     []int{3},
			k:        3,
			expected: 1,
		},
		{
			name:     "Single Element No Match",
			nums:     []int{1},
			k:        0,
			expected: 0,
		},
		{
			name:     "Multiple Zeros",
			nums:     []int{0, 0, 0},
			k:        0,
			expected: 6, // [0], [0], [0], [0,0], [0,0], [0,0,0]
		},
		{
			name:     "K is Negative",
			nums:     []int{-1, -1, 1},
			k:        -2,
			expected: 1, // [-1, -1]
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := subarraySum(tc.nums, tc.k)
			assert.Equal(t, tc.expected, result)
		})
	}
}
