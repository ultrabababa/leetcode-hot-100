package binary_tree_level_order_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/tree/common"
)

func TestLevelOrder(t *testing.T) {
	type testCase struct {
		name     string
		input    []interface{}
		expected [][]int
	}

	tests := []testCase{
		{
			name: "Example 1",
			//      3
			//    /   \
			//   9    20
			//       /  \
			//      15   7
			input:    []interface{}{3, 9, 20, nil, nil, 15, 7},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "Example 2",
			input:    []interface{}{1},
			expected: [][]int{{1}},
		},
		{
			name:     "Empty",
			input:    []interface{}{},
			expected: [][]int{}, // 注意：通常空树返回空切片
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			assert.Equal(t, tc.expected, levelOrder(root))
		})
	}
}
