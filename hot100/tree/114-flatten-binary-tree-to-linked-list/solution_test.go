package flatten_binary_tree_to_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/tree/common"
)

func TestFlatten(t *testing.T) {
	type testCase struct {
		name     string
		input    []interface{}
		expected []int // 展开后的链表值序列
	}

	tests := []testCase{
		{
			name: "Example 1",
			//      1
			//    /   \
			//   2     5
			//  / \     \
			// 3   4     6
			input:    []interface{}{1, 2, 5, 3, 4, nil, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Example 2: Empty",
			input:    []interface{}{},
			expected: []int{},
		},
		{
			name:     "Example 3: Single Node",
			input:    []interface{}{0},
			expected: []int{0},
		},
		{
			name:     "No Left Child",
			input:    []interface{}{1, nil, 2, nil, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			flatten(root)

			// 验证展开后的结果
			var result []int
			cur := root
			for cur != nil {
				// 展开后左子树必须全部为 nil
				assert.Nil(t, cur.Left, "Left child should be nil after flattening")
				result = append(result, cur.Val)
				cur = cur.Right
			}

			// 空树时，result 为 nil，需要特殊处理断言
			if len(tc.input) == 0 {
				assert.Empty(t, result)
			} else {
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
