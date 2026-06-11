package generate_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	type testCase struct {
		name     string
		n        int
		expected []string
	}

	tests := []testCase{
		{
			name:     "Example 1: n = 3",
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "Example 2: n = 1",
			n:        1,
			expected: []string{"()"},
		},
		{
			name:     "Example 3: n = 2",
			n:        2,
			expected: []string{"(())", "()()"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.expected, generateParenthesis(tc.n))
		})
	}
}
