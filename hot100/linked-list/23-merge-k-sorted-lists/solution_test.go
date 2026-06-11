package merge_k_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/linked-list/common"
)

func TestMergeKLists(t *testing.T) {
	// 辅助函数：将二维数组转换为链表数组
	createLists := func(data [][]int) []*ListNode {
		lists := make([]*ListNode, len(data))
		for i, nums := range data {
			lists[i] = CreateList(nums)
		}
		return lists
	}

	type testCase struct {
		name     string
		input    [][]int
		expected []int
	}

	tests := []testCase{
		{
			name: "Example 1",
			input: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "Example 2 - Empty input",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name:     "Example 3 - List of empty lists",
			input:    [][]int{{}, {}},
			expected: []int{},
		},
		{
			name: "Different lengths",
			input: [][]int{
				{1},
				{0, 2},
				{1, 2, 3},
			},
			expected: []int{0, 1, 1, 2, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lists := createLists(tc.input)
			mergedHead := mergeKLists(lists)
			result := ListToArray(mergedHead)
			assert.Equal(t, tc.expected, result)
		})
	}
}
