package binarysearch

import (
	"testing"
)

// 函数名必须以 Test 开头，参数必须是 t *testing.T
func TestSearch(t *testing.T) {
	// 定义一个测试用例结构体，包含输入和期望输出
	type testCase struct {
		nums     []int
		target   int
		expected int
	}

	// 这里列出你想测的数据（Table-Driven Tests 是 Go 的特色）
	tests := []testCase{
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4, // 9 的下标是 4
		},
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1, // 2 不在数组中
		},
	}

	// 遍历所有测试用例
	for _, tc := range tests {
		result := Search(tc.nums, tc.target)
		if result != tc.expected {
			t.Errorf("Search(%v, %d) = %d; want %d", tc.nums, tc.target, result, tc.expected)
		}
	}
}
