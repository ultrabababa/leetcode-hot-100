package diameter_of_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/tree/common"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	type testCase struct {
		name     string
		input    []interface{}
		expected int
	}

	tests := []testCase{
		{
			name: "Example 1",
			//      1
			//    /   \
			//   2     3
			//  / \
			// 4   5
			// 路径：4-2-1-3 或 5-2-1-3，长度为 3
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: 3,
		},
		{
			name: "Example 2",
			//   1
			//  /
			// 2
			input:    []interface{}{1, 2},
			expected: 1,
		},
		{
			name: "Path not through root",
			//        1
			//       /
			//      2
			//     / \
			//    3   4
			//   /     \
			//  5       6
			// 最长路径在左子树内部：5-3-2-4-6，长度 4
			input:    []interface{}{1, 2, nil, 3, 4, 5, nil, nil, 6},
			expected: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			assert.Equal(t, tc.expected, diameterOfBinaryTree(root))
		})
	}
}
