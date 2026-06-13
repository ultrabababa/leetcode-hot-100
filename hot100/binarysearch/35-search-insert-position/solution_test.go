package main

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	// 定义测试用例的结构体
	type testCase struct {
		name     string // 测试用例名称
		nums     []int  // 排序数组
		target   int    // 目标值
		expected int    // 期望的索引
	}

	// 初始化测试用例集合
	testCases := []testCase{
		{
			name:     "示例 1：目标值在数组中",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "示例 2：目标值不在数组中，插入中间",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "示例 3：目标值不在数组中，插入尾部",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "边界测试：目标值不在数组中，插入头部",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "边界测试：单元素数组，命中目标",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "边界测试：包含负数，插入中间",
			nums:     []int{-10, -5, 0, 3, 6},
			target:   -7,
			expected: 1,
		},
	}

	// 遍历并执行所有测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := searchInsert(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("失败 - [%s]: nums=%v, target=%d. 期望输出: %d, 实际输出: %d",
					tc.name, tc.nums, tc.target, tc.expected, result)
			}
		})
	}
}
