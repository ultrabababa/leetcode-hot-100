package palindrome_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 使用点导入我们自己写的公共包
	. "leetcode/hot100/linked-list/common"
)

func TestIsPalindrome(t *testing.T) {
	type testCase struct {
		name     string
		input    []int
		expected bool
	}

	tests := []testCase{
		{
			name:     "Example 1 - Even palindrome",
			input:    []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "Example 2 - Not palindrome",
			input:    []int{1, 2},
			expected: false,
		},
		{
			name:     "Odd palindrome",
			input:    []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := CreateList(tc.input)
			result := isPalindrome(head)
			assert.Equal(t, tc.expected, result)
		})
	}
}
