package convert_sorted_array_to_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedArrayToBST(t *testing.T) {
	type testCase struct {
		name  string
		input []int
		// 这里的 expected 我们只验证一种情况。
		// 实际上 [0,-10,5,null,-3,null,9] 也是对的，取决于偶数长度时取左中还是右中。
		// 我们的算法通常取 len/2，即右中（或者说中间偏后）。
		expected []interface{}
	}

	tests := []testCase{
		{
			name:  "Example 1",
			input: []int{-10, -3, 0, 5, 9},
			// 期望输出是层序遍历的结果，用于验证
			// 根是 0
			// 左是 -3 (左-10), 右是 9 (左5)
			expected: []interface{}{0, -3, 9, -10, nil, 5},
		},
		{
			name:     "Example 2",
			input:    []int{1, 3},
			expected: []interface{}{3, 1}, // 根 3，左 1
		},
		{
			name:     "Single Element",
			input:    []int{1},
			expected: []interface{}{1},
		},
		{
			name:     "Empty",
			input:    []int{},
			expected: []interface{}{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := sortedArrayToBST(tc.input)

			// 验证方法：
			// 1. 验证是否是二叉搜索树 (中序遍历必须有序)
			// 2. 验证是否平衡 (最大深度差不超过1)
			// 3. 简单验证：直接对比层序遍历结果 (TreeToSlice 是假设你 common 包里有这个辅助函数)
			// 如果没有 TreeToSlice，你可以简单写一个，或者只断言 root.Val

			// 这里假设你已经按照之前的建议在 common 里实现了 TreeToSlice
			// 也可以用 assert.Equal(t, tc.expected, TreeToSlice(root))

			// 最简单的非严格验证：
			if len(tc.input) > 0 {
				assert.Equal(t, tc.expected[0], root.Val)
			} else {
				assert.Nil(t, root)
			}
		})
	}
}
