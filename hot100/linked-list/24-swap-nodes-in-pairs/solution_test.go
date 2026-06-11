package swap_nodes_in_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入我们天下无敌的 common 包
	. "leetcode/hot100/linked-list/common"
)

func TestSwapPairs(t *testing.T) {
	type testCase struct {
		name     string
		head     []int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1 - Even number of nodes",
			head:     []int{1, 2, 3, 4},
			expected: []int{2, 1, 4, 3},
		},
		{
			name:     "Odd number of nodes",
			head:     []int{1, 2, 3},
			expected: []int{2, 1, 3}, // 3 保持原样
		},
		{
			name:     "Example 2 - Empty list",
			head:     []int{},
			expected: []int{},
		},
		{
			name:     "Example 3 - Single element",
			head:     []int{1},
			expected: []int{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := CreateList(tc.head)
			resultList := swapPairs(head)
			resultArray := ListToArray(resultList)
			assert.Equal(t, tc.expected, resultArray)
		})
	}
}
