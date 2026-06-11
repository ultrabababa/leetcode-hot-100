package combination_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	type testCase struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}

	tests := []testCase{
		{
			name:       "Example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			// 我们代码中会先排序，所以先探索 2，再 3...
			expected: [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Example 3: Target too small",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{}, // nil 或者 空数组都可以，只要内容为空
		},
		{
			name:       "Single element exact match",
			candidates: []int{7},
			target:     7,
			expected:   [][]int{{7}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := combinationSum(tc.candidates, tc.target)
			if len(tc.expected) == 0 {
				assert.Empty(t, res)
			} else {
				assert.Equal(t, tc.expected, res)
			}
		})
	}
}
