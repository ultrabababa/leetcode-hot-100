package invert_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/tree/common"
)

func TestInvertTree(t *testing.T) {
	type testCase struct {
		name     string
		input    []interface{}
		expected []int // 这里的 expected 我们用层序遍历的 int 数组来对比
	}

	tests := []testCase{
		{
			name:  "Example 1",
			input: []interface{}{4, 2, 7, 1, 3, 6, 9},
			// 翻转后：
			//      4
			//    /   \
			//   7     2
			//  / \   / \
			// 9   6 3   1
			expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "Example 2",
			input:    []interface{}{2, 1, 3},
			expected: []int{2, 3, 1},
		},
		{
			name:     "Empty",
			input:    []interface{}{},
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			invertedRoot := invertTree(root)

			// 为了验证结果，我们需要把树转回数组
			// 复用 common 包里的 TreeToSlice (如果你在 common 里写了这个辅助函数)
			// 或者简单的手动对比，这里假设你有 TreeToSlice
			result := TreeToSlice(invertedRoot)
			assert.Equal(t, tc.expected, result)
		})
	}
}
