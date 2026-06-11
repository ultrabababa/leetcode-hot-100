package maximum_depth_of_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包，复用 SliceToTree
	. "leetcode/hot100/tree/common"
)

func TestMaxDepth(t *testing.T) {
	type testCase struct {
		name     string
		input    []interface{}
		expected int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			input:    []interface{}{3, 9, 20, nil, nil, 15, 7},
			expected: 3,
		},
		{
			name:     "Example 2",
			input:    []interface{}{1, nil, 2},
			expected: 2,
		},
		{
			name:     "Empty Tree",
			input:    []interface{}{},
			expected: 0,
		},
		{
			name:     "Single Node",
			input:    []interface{}{1},
			expected: 1,
		},
		{
			name:     "Left Skewed",
			input:    []interface{}{1, 2, nil, 3}, // 1->2->3
			expected: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			assert.Equal(t, tc.expected, maxDepth(root))
		})
	}
}
