package validate_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/tree/common"
)

func TestIsValidBST(t *testing.T) {
	type testCase struct {
		name     string
		input    []interface{}
		expected bool
	}

	tests := []testCase{
		{
			name:     "Example 1: Valid",
			input:    []interface{}{2, 1, 3},
			expected: true,
		},
		{
			name: "Example 2: Invalid (Child ok, Grandchild wrong)",
			//      5
			//    /   \
			//   1     4
			//        / \
			//       3   6
			// 3 小于 4 (局部OK)，但 3 也小于 5 (根)！
			// 根据 BST 规则，5 的右子树里所有节点都必须大于 5。所以这是 false。
			input:    []interface{}{5, 1, 4, nil, nil, 3, 6},
			expected: false,
		},
		{
			name:     "Invalid: Left > Root",
			input:    []interface{}{2, 3, 3}, // 左孩子3 > 根2
			expected: false,
		},
		{
			name:     "Invalid: Equal values",
			input:    []interface{}{2, 2, 2}, // BST 定义通常是严格小于/大于
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			assert.Equal(t, tc.expected, isValidBST(root))
		})
	}
}
