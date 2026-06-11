package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type testCase struct {
		name     string
		s        string
		expected int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			s:        "abcabcbb",
			expected: 3, // "abc"
		},
		{
			name:     "Example 2 - All Same",
			s:        "bbbbb",
			expected: 1, // "b"
		},
		{
			name:     "Example 3 - Overlap",
			s:        "pwwkew",
			expected: 3, // "wke"
		},
		{
			name:     "Empty String",
			s:        "",
			expected: 0,
		},
		{
			name:     "Single Character",
			s:        " ", // 注意：空格也是字符
			expected: 1,
		},
		{
			name:     "Tricky Case - Window Jump",
			s:        "dvdf",
			expected: 3, // "vdf"。很多错解会输出 2 (dv)
		},
		{
			name:     "Start with Repeat",
			s:        "abba",
			expected: 2, // "ab" or "ba"
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.s)
			assert.Equal(t, tc.expected, result)
		})
	}
}
