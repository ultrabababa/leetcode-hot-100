package letter_combinations_of_a_phone_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	type testCase struct {
		name     string
		digits   string
		expected []string
	}

	tests := []testCase{
		{
			name:   "Example 1: '23'",
			digits: "23",
			// 2 -> a,b,c; 3 -> d,e,f
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "Example 2: '2'",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Empty string",
			digits:   "",
			expected: []string{}, // 注意：LeetCode 要求空字符串返回空数组，而不是 [""]
		},
		{
			name:   "Example with 4 letters: '79'",
			digits: "79",
			// 7 -> pqrs; 9 -> wxyz
			expected: []string{
				"pw", "px", "py", "pz",
				"qw", "qx", "qy", "qz",
				"rw", "rx", "ry", "rz",
				"sw", "sx", "sy", "sz",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.expected, letterCombinations(tc.digits))
		})
	}
}
