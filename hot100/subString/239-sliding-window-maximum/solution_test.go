package sliding_window_maximum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		k        int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "Example 2 - Single Element Window",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Example 2 - k equals length",
			nums:     []int{9, 11},
			k:        2,
			expected: []int{11},
		},
		{
			name:     "Monotonically Increasing",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 4, 5}, // 每次最大值都在窗口最右边
		},
		{
			name:     "Monotonically Decreasing",
			nums:     []int{5, 4, 3, 2, 1},
			k:        3,
			expected: []int{5, 4, 3}, // 每次最大值都在窗口最左边
		},
		{
			name:     "Negative Numbers",
			nums:     []int{-7, -8, 7, 5, 7, 1, 6, 0},
			k:        4,
			expected: []int{7, 7, 7, 7, 7},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maxSlidingWindow(tc.nums, tc.k)
			assert.Equal(t, tc.expected, result)
		})
	}
}
