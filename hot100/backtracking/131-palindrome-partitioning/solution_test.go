package palindrome_partitioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	type testCase struct {
		name     string
		s        string
		expected [][]string
	}

	tests := []testCase{
		{
			name: "Example 1: aab",
			s:    "aab",
			expected: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "Example 2: a",
			s:    "a",
			expected: [][]string{
				{"a"},
			},
		},
		{
			name: "All same characters",
			s:    "aaa",
			expected: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name: "No multiple palindromes",
			s:    "ab",
			expected: [][]string{
				{"a", "b"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.expected, partition(tc.s))
		})
	}
}
