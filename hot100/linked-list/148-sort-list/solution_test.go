package sort_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/linked-list/common"
)

func TestSortList(t *testing.T) {
	type testCase struct {
		name     string
		head     []int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1 - Random Order",
			head:     []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2 - Negative & Random",
			head:     []int{-1, 5, 3, 4, 0},
			expected: []int{-1, 0, 3, 4, 5},
		},
		{
			name:     "Example 3 - Empty List",
			head:     []int{},
			expected: []int{},
		},
		{
			name:     "Already Sorted",
			head:     []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Reverse Sorted",
			head:     []int{4, 3, 2, 1},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Duplicates",
			head:     []int{4, 1, 2, 1},
			expected: []int{1, 1, 2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := CreateList(tc.head)
			sortedHead := sortList(head)
			result := ListToArray(sortedHead)
			assert.Equal(t, tc.expected, result)
		})
	}
}
