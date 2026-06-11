package course_schedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	type testCase struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}

	tests := []testCase{
		{
			name:          "Example 1: Acyclic",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}}, // 0 -> 1
			expected:      true,
		},
		{
			name:          "Example 2: Cyclic",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}}, // 0 -> 1 -> 0
			expected:      false,
		},
		{
			name:          "No prerequisites",
			numCourses:    3,
			prerequisites: [][]int{}, // 各学各的
			expected:      true,
		},
		{
			name:       "Larger Acyclic Graph",
			numCourses: 4,
			// 0 -> 1, 0 -> 2, 1 -> 3, 2 -> 3
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      true,
		},
		{
			name:       "Larger Cyclic Graph",
			numCourses: 4,
			// 0 -> 1 -> 2 -> 3 -> 1 (环在 1-2-3-1)
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}},
			expected:      false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, canFinish(tc.numCourses, tc.prerequisites))
		})
	}
}
