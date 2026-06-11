package word_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	type testCase struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}

	tests := []testCase{
		{
			name: "Example 1: ABCCED",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			name: "Example 2: SEE",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			name: "Example 3: ABCB (Cannot reuse cell)",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
		{
			name: "Single Character Match",
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
		},
		{
			name: "Single Character Mismatch",
			board: [][]byte{
				{'A'},
			},
			word:     "B",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, exist(tc.board, tc.word))
		})
	}
}
