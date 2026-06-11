package reverse_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"

	// 测试文件里同样使用点导入
	. "leetcode/hot100/linked-list/common"
)

func TestReverseList(t *testing.T) {
	type testCase struct {
		name     string
		input    []int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Example 2 - Two elements",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Example 3 - Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 1. 利用辅助函数：数组 -> 链表
			head := CreateList(tc.input)

			// 2. 执行你的算法
			reversedHead := reverseList(head)

			// 3. 利用辅助函数：链表 -> 数组
			result := ListToArray(reversedHead)

			// 4. 直接对比数组，极其直观！
			assert.Equal(t, tc.expected, result)
		})
	}
}
