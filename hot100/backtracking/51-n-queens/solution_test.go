package n_queens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	type testCase struct {
		name     string
		n        int
		expected int // 为了简化测试，我们这里只校验解的数量
	}

	tests := []testCase{
		{
			name:     "Example 1: 4 Queens",
			n:        4,
			expected: 2, // 4皇后有两个解
		},
		{
			name:     "Example 2: 1 Queen",
			n:        1,
			expected: 1,
		},
		{
			name:     "8 Queens Classic Problem",
			n:        8,
			expected: 92, // 经典的8皇后问题有92个解
		},
		{
			name:     "9 Queens",
			n:        9,
			expected: 352,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := solveNQueens(tc.n)
			assert.Equal(t, tc.expected, len(res))

			// 简单校验一下输出格式（以 4 皇后为例）
			if tc.n == 4 {
				expectedSolutions := [][]string{
					{".Q..", "...Q", "Q...", "..Q."},
					{"..Q.", "Q...", "...Q", ".Q.."},
				}
				assert.ElementsMatch(t, expectedSolutions, res)
			}
		})
	}
}
