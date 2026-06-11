package binary_tree_maximum_path_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/tree/common"
)

func TestMaxPathSum(t *testing.T) {
	type testCase struct {
		name     string
		input    []interface{}
		expected int
	}

	tests := []testCase{
		{
			name: "Example 1",
			//   1
			//  / \
			// 2   3
			// 路径: 2 -> 1 -> 3
			input:    []interface{}{1, 2, 3},
			expected: 6,
		},
		{
			name: "Example 2",
			//    -10
			//    /  \
			//   9   20
			//      /  \
			//     15   7
			// 路径: 15 -> 20 -> 7
			input:    []interface{}{-10, 9, 20, nil, nil, 15, 7},
			expected: 42,
		},
		{
			name: "All Negatives",
			//   -3
			// 路径只能是 -3 本身，不能是 0（题目要求至少包含一个节点）
			input:    []interface{}{-3},
			expected: -3,
		},
		{
			name: "Negative Child Abandonment",
			//    2
			//   / \
			// -1  -2
			// 最优路径就是 2 节点自己，抛弃两个负数子节点
			input:    []interface{}{2, -1, -2},
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			assert.Equal(t, tc.expected, maxPathSum(root))
		})
	}
}
