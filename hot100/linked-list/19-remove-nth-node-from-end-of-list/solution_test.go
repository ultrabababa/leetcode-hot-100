package remove_nth_node_from_end_of_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入我们天下无敌的 common 包
	. "leetcode/hot100/linked-list/common"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type testCase struct {
		name     string
		head     []int
		n        int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1 - Normal case",
			head:     []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5}, // 删除了倒数第 2 个节点 (4)
		},
		{
			name:     "Example 2 - Single element",
			head:     []int{1},
			n:        1,
			expected: []int{}, // 删除了唯一的节点
		},
		{
			name:     "Example 3 - Two elements, remove last",
			head:     []int{1, 2},
			n:        1,
			expected: []int{1}, // 删除了倒数第 1 个节点 (2)
		},
		{
			name:     "Remove the head node",
			head:     []int{1, 2, 3},
			n:        3,
			expected: []int{2, 3}, // 删除了倒数第 3 个节点，也就是头节点 (1)
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := CreateList(tc.head)
			resultList := removeNthFromEnd(head, tc.n)
			resultArray := ListToArray(resultList)
			assert.Equal(t, tc.expected, resultArray)
		})
	}
}
