package merge_intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	type testCase struct {
		name      string
		intervals [][]int
		expected  [][]int
	}

	tests := []testCase{
		{
			name:      "Example 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Example 2 - Touching Intervals",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "Example 3 - Unsorted Input",
			intervals: [][]int{{4, 7}, {1, 4}},
			expected:  [][]int{{1, 7}}, // 排序后变成 [1,4],[4,7] -> [1,7]
		},
		{
			name:      "Fully Contained",
			intervals: [][]int{{1, 10}, {2, 6}, {3, 5}},
			expected:  [][]int{{1, 10}}, // 后面的都在第一个肚子里
		},
		{
			name:      "No Overlap",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected:  [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:      "Single Interval",
			intervals: [][]int{{1, 4}},
			expected:  [][]int{{1, 4}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := merge(tc.intervals)
			assert.Equal(t, tc.expected, result)
		})
	}
}
