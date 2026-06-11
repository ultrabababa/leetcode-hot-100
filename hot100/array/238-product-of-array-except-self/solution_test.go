package product_of_array_except_self

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2 - Contains Zero",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Two Zeros",
			nums:     []int{0, 4, 0},
			expected: []int{0, 0, 0}, // 任何位置的“其余元素”里都至少包含一个0
		},
		{
			name:     "Negative Numbers",
			nums:     []int{-1, -1},
			expected: []int{-1, -1},
		},
		{
			name:     "Mixed Signs",
			nums:     []int{2, -3, 4},
			expected: []int{-12, 8, -6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := productExceptSelf(tc.nums)
			assert.Equal(t, tc.expected, result)
		})
	}
}
