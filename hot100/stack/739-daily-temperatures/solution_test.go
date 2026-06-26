package main

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		expected     []int
	}{
		{
			name:         "Example 1 - Mixed",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "Example 2 - Strictly Increasing",
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		{
			name:         "Example 3 - Increasing",
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
		{
			name:         "Strictly Decreasing",
			temperatures: []int{90, 80, 70, 60},
			expected:     []int{0, 0, 0, 0},
		},
		{
			name:         "All Same",
			temperatures: []int{50, 50, 50},
			expected:     []int{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("dailyTemperatures(%v) = %v, expected %v", tt.temperatures, got, tt.expected)
			}
		})
	}
}
