package find_all_anagrams_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	type testCase struct {
		name     string
		s        string
		p        string
		expected []int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},
		{
			name:     "Example 2 - Overlap",
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
		{
			name:     "No Solution",
			s:        "hello",
			p:        "world",
			expected: []int{}, // 空切片
		},
		{
			name:     "S is shorter than P",
			s:        "ab",
			p:        "abc",
			expected: []int{},
		},
		{
			name:     "Exact Match",
			s:        "abc",
			p:        "abc",
			expected: []int{0},
		},
		{
			name:     "Repeated Characters",
			s:        "aaabbb",
			p:        "aabb",
			expected: []int{1}, // index 1: "aabb"
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := findAnagrams(tc.s, tc.p)

			// 结果必须匹配（顺序其实是固定的从小到大，直接 Equal 即可）
			// 如果你的算法产生的顺序不确定，可以用 assert.ElementsMatch
			assert.Equal(t, tc.expected, result)
		})
	}
}
