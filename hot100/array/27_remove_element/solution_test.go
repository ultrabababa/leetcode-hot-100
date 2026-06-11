package removeelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	type testCase struct {
		nums     []int
		val      int
		expected int    // 期望的新长度
		name     string // 给测试用例起个名字，报错时好找
	}

	tests := []testCase{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
		},
		{
			name:     "Empty Array",
			nums:     []int{},
			val:      0,
			expected: 0,
		},
		{
			name:     "Remove All",
			nums:     []int{1, 1, 1},
			val:      1,
			expected: 0,
		},
		{
			name:     "Remove None",
			nums:     []int{1, 2, 3},
			val:      5,
			expected: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 1. 调用你的函数
			// 注意：因为是原地修改，传切片进去，函数内修改底层数组，外层也能看到变化
			k := RemoveElement(tc.nums, tc.val)

			// 2. 检查返回值（长度）是否正确
			assert.Equal(t, tc.expected, k, "New length should be correct")

			// 3. 关键步骤：检查数组的前 k 个元素里，是否真的没有 val 了
			// 我们只关心 nums[:k]，后面的元素是什么都不重要
			resultSlice := tc.nums[:k]
			for _, x := range resultSlice {
				assert.NotEqual(t, tc.val, x, "Result slice should not contain the removed value")
			}
		})
	}
}
