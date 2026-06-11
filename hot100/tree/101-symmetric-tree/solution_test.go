package symmetric_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 注意：根据你的 import 路径，这里引入 common 包
	. "leetcode/hot100/tree/common"
)

func TestIsSymmetric(t *testing.T) {
	type testCase struct {
		name     string
		input    []interface{}
		expected bool
	}

	tests := []testCase{
		{
			name: "Example 1: Symmetric",
			//      1
			//    /   \
			//   2     2
			//  / \   / \
			// 3   4 4   3
			input:    []interface{}{1, 2, 2, 3, 4, 4, 3},
			expected: true,
		},
		{
			name: "Example 2: Asymmetric (Structure)",
			//      1
			//    /   \
			//   2     2
			//    \     \
			//     3     3
			input:    []interface{}{1, 2, 2, nil, 3, nil, 3},
			expected: false,
		},
		{
			name: "Asymmetric (Value)",
			//      1
			//    /   \
			//   2     2
			//  /       \
			// 3         4  (3 != 4)
			input:    []interface{}{1, 2, 2, 3, nil, nil, 4},
			expected: false,
		},
		{
			name:     "Single Node",
			input:    []interface{}{1},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			// 测试递归
			assert.Equal(t, tc.expected, isSymmetric(root))
			// 测试迭代
			assert.Equal(t, tc.expected, isSymmetricIterative(root))
		})
	}
}
