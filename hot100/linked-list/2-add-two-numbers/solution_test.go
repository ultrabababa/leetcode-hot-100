package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入 common 包
	. "leetcode/hot100/linked-list/common"
)

func TestAddTwoNumbers(t *testing.T) {
	type testCase struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			l1:       []int{2, 4, 3}, // 342
			l2:       []int{5, 6, 4}, // 465
			expected: []int{7, 0, 8}, // 807
		},
		{
			name:     "Example 2 - Zeros",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "Example 3 - Different Lengths & Carry",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name:     "Simple Carry at End",
			l1:       []int{5},
			l2:       []int{5},
			expected: []int{0, 1}, // 5+5=10 -> [0, 1]
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list1 := CreateList(tc.l1)
			list2 := CreateList(tc.l2)

			resultList := addTwoNumbers(list1, list2)

			resultArray := ListToArray(resultList)
			assert.Equal(t, tc.expected, resultArray)
		})
	}
}
