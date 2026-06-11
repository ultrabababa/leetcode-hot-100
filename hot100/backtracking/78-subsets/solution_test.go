package subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	type testCase struct {
		name     string
		input    []int
		expected [][]int
	}

	tests := []testCase{
		{
			name:  "Example 1: [1,2,3]",
			input: []int{1, 2, 3},
			// 按照 DFS 的访问顺序：
			// [] -> [1] -> [1,2] -> [1,2,3] -> [1,3] -> [2] -> [2,3] -> [3]
			expected: [][]int{
				{},
				{1}, {1, 2}, {1, 2, 3}, {1, 3},
				{2}, {2, 3},
				{3},
			},
		},
		{
			name:  "Example 2: [0]",
			input: []int{0},
			expected: [][]int{
				{}, {0},
			},
		},
		{
			name:  "Empty Array",
			input: []int{},
			expected: [][]int{
				{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 子集问题其实不需要 ElementsMatch，因为我们严格按照字典序前缀生成
			assert.Equal(t, tc.expected, subsets(tc.input))
		})
	}
}
