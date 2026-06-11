package minimum_window_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	type testCase struct {
		name     string
		s        string
		t        string
		expected string
	}

	tests := []testCase{
		{
			name:     "Example 1",
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
		},
		{
			name:     "Example 2 - Exact Match",
			s:        "a",
			t:        "a",
			expected: "a",
		},
		{
			name:     "Example 3 - Insufficient Count",
			s:        "a",
			t:        "aa",
			expected: "", // s 只有一个 a，不够
		},
		{
			name:     "Target longer than Source",
			s:        "a",
			t:        "ab",
			expected: "",
		},
		{
			name:     "Duplicates in Target",
			s:        "aaflslflsldkalskaaa",
			t:        "aaa",
			expected: "aaa", // 必须包含3个a
		},
		{
			name:     "Substr at end",
			s:        "ab",
			t:        "b",
			expected: "b",
		},
		{
			name:     "Irrelevant chars mixed",
			s:        "bba",
			t:        "ab",
			expected: "ba", // 包含 b, a，且最短
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := minWindow(tc.s, tc.t)
			assert.Equal(t, tc.expected, result)
		})
	}
}
