package reverse_nodes_in_k_group

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入公共包
	. "leetcode/hot100/linked-list/common"
)

func TestReverseKGroup(t *testing.T) {
	type testCase struct {
		name     string
		head     []int
		k        int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1 - Partial reverse at end",
			head:     []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{2, 1, 4, 3, 5}, // 5 保持原样
		},
		{
			name:     "Example 2 - Exact multiple",
			head:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 2, 1, 4, 5}, // 4, 5 保持原样
		},
		{
			name:     "k = 1, no change",
			head:     []int{1, 2, 3},
			k:        1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "k equals length",
			head:     []int{1, 2, 3},
			k:        3,
			expected: []int{3, 2, 1},
		},
		{
			name:     "Empty list",
			head:     []int{},
			k:        2,
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := CreateList(tc.head)
			resultList := reverseKGroup(head, tc.k)
			resultArray := ListToArray(resultList)
			assert.Equal(t, tc.expected, resultArray)
		})
	}
}
