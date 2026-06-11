package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	type testCase struct {
		name     string
		input    []int
		expected [][]int
	}

	tests := []testCase{
		{
			name:  "Example 1: [1,2,3]",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name:     "Example 2: [0,1]",
			input:    []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			name:     "Example 3: [1]",
			input:    []int{1},
			expected: [][]int{{1}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.expected, permute(tc.input))
		})
	}
}
