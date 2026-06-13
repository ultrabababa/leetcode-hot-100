package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 2d -> 1d index range
	left, right := 0, m*n-1

	for left <= right {
		mid := (left + right) / 2
		// 1d index map to (x,y)
		rmid, cmid := mid/n, mid%n
		if matrix[rmid][cmid] == target {
			return true
		} else if matrix[rmid][cmid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
