package binary_tree_inorder_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入你刚写的 common 包
	. "leetcode/hot100/tree/common"
)

func TestInorderTraversal(t *testing.T) {
	type testCase struct {
		name     string
		input    []interface{} // 使用 interface{} 兼容 nil
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1: [1,null,2,3]",
			input:    []interface{}{1, nil, 2, 3}, // 直接照抄 LeetCode 的输入
			expected: []int{1, 3, 2},
		},
		{
			name:     "Example 2: Empty",
			input:    []interface{}{},
			expected: []int{},
		},
		{
			name:     "Example 3: Single Node",
			input:    []interface{}{1},
			expected: []int{1},
		},
		{
			name: "Full Tree",
			//      1
			//    /   \
			//   2     3
			//  / \
			// 4   5
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []int{4, 2, 5, 1, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 一行代码构建树，爽不爽？
			root := SliceToTree(tc.input)

			assert.Equal(t, tc.expected, inorderTraversal(root))
		})
	}
}
