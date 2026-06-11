package container_with_most_water

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxRes := 0
	for left < right {
		width := right - left
		currArea := 0
		if height[left] <= height[right] {
			currArea = width * height[left]
			left++
		} else {
			currArea = width * height[right]
			right--
		}
		if currArea > maxRes {
			maxRes = currArea
		}
	}
	return maxRes
}
