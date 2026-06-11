package threesum

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		expected [][]int
	}

	tests := []testCase{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name:     "Example 2 - No Solution",
			nums:     []int{0, 1, 1},
			expected: [][]int{}, // 空结果
		},
		{
			name:     "Example 3 - All Zeros",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "More Zeros",
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}}, // 哪怕有4个0，去重后也只能有一组
		},
		{
			name:     "Mixed Duplicates",
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := threeSum(tc.nums)

			// 必须把期望值和结果都排个序，才能比较
			assert.Equal(t, sortResult(tc.expected), sortResult(result))
		})
	}
}

// 辅助函数：将二维切片排序（标准化），逻辑类似之前的字母异位词
func sortResult(res [][]int) [][]int {
	// 1. 深拷贝，防止修改原数据
	newRes := make([][]int, len(res))
	for i, v := range res {
		// 复制并对内部排序 (例如 [-1, 2, -1] -> [-1, -1, 2])
		temp := make([]int, len(v))
		copy(temp, v)
		sort.Ints(temp)
		newRes[i] = temp
	}

	// 2. 对外部排序 (例如 [[-1,0,1], [-1,-1,2]] -> [[-1,-1,2], [-1,0,1]])
	sort.Slice(newRes, func(i, j int) bool {
		// 逐个元素比较
		for k := 0; k < 3; k++ {
			if newRes[i][k] != newRes[j][k] {
				return newRes[i][k] < newRes[j][k]
			}
		}
		return false
	})
	return newRes
}
