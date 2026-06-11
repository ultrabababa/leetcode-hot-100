package groupanagrams

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	type testCase struct {
		name     string
		strs     []string
		expected [][]string
	}

	tests := []testCase{
		{
			name: "Example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
		{
			name:     "Empty String",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Single Character",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "No Anagrams",
			strs:     []string{"hello", "world", "code"},
			expected: [][]string{{"hello"}, {"world"}, {"code"}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := groupAnagrams(tc.strs)

			// 1. 简单检查：组的数量必须一致
			assert.Equal(t, len(tc.expected), len(result), "Should have same number of groups")

			// 2. 核心检查：因为题目说“可以按任意顺序返回”，
			// 所以我们在比较前，必须把 expected 和 result 都“标准化”（排序）
			// 否则虽然答案是对的，但顺序不一样会导致测试失败
			sortedExpected := sortGroups(tc.expected)
			sortedResult := sortGroups(result)

			assert.Equal(t, sortedExpected, sortedResult, "Groups content should match")
		})
	}
}

// sortGroups 是一个辅助函数，用于将结果“标准化”，忽略顺序差异
// 它做两件事：
// 1. 把每个小组内部的字符串排序 (比如 ["tea", "eat"] -> ["eat", "tea"])
// 2. 把所有小组根据第一个元素排序 (比如 [["bat"], ["eat", "tea"]] -> 按照 b, e 的顺序排)
func sortGroups(groups [][]string) [][]string {
	// 深拷贝一份，以免修改原始数据
	// 注意：在测试代码中，简单的深拷贝通常足够。为了代码简洁，这里我们新建切片。
	newGroups := make([][]string, len(groups))
	for i, group := range groups {
		// 复制内部 slice
		newGroup := make([]string, len(group))
		copy(newGroup, group)
		// 1. 组内排序
		sort.Strings(newGroup)
		newGroups[i] = newGroup
	}

	// 2. 组间排序
	// 我们根据每个组的第一个元素（也就是最小的那个字符串）来决定组的顺序
	sort.Slice(newGroups, func(i, j int) bool {
		// 防御性编程：如果有空组，把它排前面
		if len(newGroups[i]) == 0 {
			return true
		}
		if len(newGroups[j]) == 0 {
			return false
		}

		// 比较两个组的第一个元素
		return newGroups[i][0] < newGroups[j][0]
	})

	return newGroups
}
