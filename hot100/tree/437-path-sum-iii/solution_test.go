package path_sum_iii

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/tree/common"
)

func TestPathSum(t *testing.T) {
	type testCase struct {
		name      string
		input     []interface{}
		targetSum int
		expected  int
	}

	tests := []testCase{
		{
			name: "Example 1",
			//       10
			//      /  \
			//     5   -3
			//    / \    \
			//   3   2   11
			//  / \   \
			// 3  -2   1
			input:     []interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1},
			targetSum: 8,
			// 路径1: 5 -> 3
			// 路径2: 5 -> 2 -> 1
			// 路径3: -3 -> 11
			expected: 3,
		},
		{
			name:      "Example 2",
			input:     []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1},
			targetSum: 22,
			expected:  3,
		},
		{
			name:      "Empty Tree",
			input:     []interface{}{},
			targetSum: 0,
			expected:  0,
		},
		{
			name:      "Single Node Match",
			input:     []interface{}{1},
			targetSum: 1,
			expected:  1,
		},
		{
			name:      "Zero Sums",
			input:     []interface{}{0, 0, 0}, // 存在多条和为0的路径
			targetSum: 0,
			expected:  5, // [0]根, [0]左, [0]右, [0,0]根左, [0,0]根右
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := SliceToTree(tc.input)
			assert.Equal(t, tc.expected, pathSum(root, tc.targetSum))
		})
	}
}
