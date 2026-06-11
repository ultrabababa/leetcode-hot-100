package lowest_common_ancestor_of_a_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/tree/common"
)

func TestLowestCommonAncestor(t *testing.T) {
	// 辅助函数：根据节点值在树中寻找真实的节点指针
	var findNode func(root *TreeNode, val int) *TreeNode
	findNode = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == val {
			return root
		}
		left := findNode(root.Left, val)
		if left != nil {
			return left
		}
		return findNode(root.Right, val)
	}

	type testCase struct {
		name     string
		input    []interface{}
		pVal     int
		qVal     int
		expected int
	}

	tests := []testCase{
		{
			name: "Example 1: p and q on different sides",
			//      3
			//    /   \
			//   5     1
			//  / \   / \
			// 6   2 0   8
			//    / \
			//   7   4
			input:    []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
			pVal:     5,
			qVal:     1,
			expected: 3,
		},
		{
			name: "Example 2: p is ancestor of q",
			// p=5, q=4. 4 is a descendant of 5, so LCA is 5.
			input:    []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
			pVal:     5,
			qVal:     4,
			expected: 5,
		},
		{
			name:     "Example 3: Small tree",
			input:    []interface{}{1, 2},
			pVal:     1,
			qVal:     2,
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			p := findNode(root, tc.pVal)
			q := findNode(root, tc.qVal)

			lca := lowestCommonAncestor(root, p, q)

			// 验证最近公共祖先的值
			assert.NotNil(t, lca)
			assert.Equal(t, tc.expected, lca.Val)
		})
	}
}
