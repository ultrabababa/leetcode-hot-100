package merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		last := res[len(res)-1]
		if curr[0] <= last[1] {
			// 重叠
			last[1] = max(last[1], curr[1])
		} else {
			res = append(res, curr)
		}
	}

	return res
}
