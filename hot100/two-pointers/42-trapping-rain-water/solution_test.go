package trapping_rain_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	type testCase struct {
		name     string
		height   []int
		expected int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			name:     "Example 2",
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
		{
			name:     "Empty Array",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "Flat Land",
			height:   []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Ascending Stairs",
			height:   []int{0, 1, 2, 3, 4},
			expected: 0, // 递增，存不住水
		},
		{
			name:     "Descending Stairs",
			height:   []int{4, 3, 2, 1, 0},
			expected: 0, // 递减，存不住水
		},
		{
			name:     "Mountain Shape",
			height:   []int{0, 2, 4, 2, 0},
			expected: 0, // 凸字形，存不住水
		},
		{
			name:     "V Shape",
			height:   []int{4, 0, 4},
			expected: 4, // 凹字形，中间存 4
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := trap(tc.height)
			assert.Equal(t, tc.expected, result)
		})
	}
}
