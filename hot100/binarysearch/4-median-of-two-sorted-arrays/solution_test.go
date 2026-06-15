package main

import (
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "Example 1: Total length is odd",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.00000,
		},
		{
			name:  "Example 2: Total length is even",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.50000,
		},
		{
			name:  "Edge Case: First array is empty",
			nums1: []int{},
			nums2: []int{1, 2, 3, 4, 5},
			want:  3.00000,
		},
		{
			name:  "Edge Case: Second array is empty",
			nums1: []int{2, 2, 4, 4},
			nums2: []int{},
			want:  3.00000,
		},
		{
			name:  "Edge Case: Arrays with completely disjoint ranges",
			nums1: []int{1, 2, 3},
			nums2: []int{8, 9, 10},
			want:  5.50000,
		},
		{
			name:  "Edge Case: Arrays of different lengths",
			nums1: []int{1, 3, 8, 9, 15},
			nums2: []int{7, 11, 18, 19, 21, 25},
			want:  11.00000,
		},
	}

	const epsilon = 1e-5 // 处理浮点数比较的容差

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if math.Abs(got-tt.want) > epsilon {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
