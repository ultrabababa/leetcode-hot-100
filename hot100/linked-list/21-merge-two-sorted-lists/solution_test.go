package merge_two_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入你的 common 包
	. "leetcode/hot100/linked-list/common"
)

func TestMergeTwoLists(t *testing.T) {
	type testCase struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			l1:       []int{1, 2, 4},
			l2:       []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "Example 2 - Both empty",
			l1:       []int{},
			l2:       []int{},
			expected: []int{},
		},
		{
			name:     "Example 3 - One empty",
			l1:       []int{},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "Different lengths",
			l1:       []int{1, 2},
			l2:       []int{3, 4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "One list much larger values",
			l1:       []int{10, 20},
			l2:       []int{1, 2},
			expected: []int{1, 2, 10, 20},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 1. 造链表
			list1 := CreateList(tc.l1)
			list2 := CreateList(tc.l2)

			// 2. 运行算法
			mergedHead := mergeTwoLists(list1, list2)

			// 3. 转回数组并断言
			result := ListToArray(mergedHead)
			assert.Equal(t, tc.expected, result)
		})
	}
}
