package rotting_oranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting(t *testing.T) {
	type testCase struct {
		name     string
		input    [][]int
		expected int
	}

	tests := []testCase{
		{
			name: "Example 1: Normal rotting",
			input: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Example 2: Unreachable orange",
			// 左下角的 1 被 0 挡住了，永远不会腐烂
			input: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			name: "Example 3: Already all rotten/empty",
			// 一开始就没有新鲜橘子，需要 0 分钟
			input: [][]int{
				{0, 2},
			},
			expected: 0,
		},
		{
			name: "Multiple sources at start",
			// 两个腐烂源头同时开始
			input: [][]int{
				{2, 1, 1},
				{1, 1, 1},
				{0, 1, 2},
			},
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 深度拷贝二维切片，防止测试用例互相污染（因为 BFS 会修改原数组）
			gridCopy := make([][]int, len(tc.input))
			for i := range tc.input {
				gridCopy[i] = make([]int, len(tc.input[i]))
				copy(gridCopy[i], tc.input[i])
			}
			assert.Equal(t, tc.expected, orangesRotting(gridCopy))
		})
	}
}
