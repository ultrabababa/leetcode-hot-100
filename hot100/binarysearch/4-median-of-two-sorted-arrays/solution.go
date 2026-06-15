// hot100/binarysearch/4-median-of-two-sorted-arrays/solution.go

package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		// 确保nums1更短，以保证复杂度为o(log(min(m,n)))
		// 且保证j不会导致数组越界
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m // 这是为了二分查找切分线i的位置，切分线有可能在最右边m，表示整个数组全是整体的左半部份
	total := (m + n + 1) / 2

	// 核心思想是用二分查找来找到一个满足条件的i,j分割点
	for left <= right {
		// i 是nums1的分界index，i的左半部份(不含i)放在总的左半部份
		// j 是nums2的分界index，i + j = (m+n+1)/2
		// 这样保证了整体的左右两部份元素数量满足条件
		i := left + (right-left)/2
		j := total - i

		// 分割完之后，整体的左半部份是nums1[0]到nums1[i-1] + nums2[0]到nums2[j-1]
		// 整体的右半部份是nums1[i]到nums1[m-1] + nums2[j]到nums2[n-1]
		// 因为都各自升序，所以只需满足条件1 nums1[i-1] <= nums2[j] 且条件2 nums2[j-1]<=nums1[i]
		// 同时还要处理边界条件
		nums1LeftMax := math.MinInt // 如果i=0，表示左侧没有任何数，条件1天然成立
		if i > 0 {
			nums1LeftMax = nums1[i-1]
		}

		nums2LeftMax := math.MinInt // j 同理
		if j > 0 {
			nums2LeftMax = nums2[j-1]
		}

		nums1RightMin := math.MaxInt // 如果i=m, 表示数组1没有任何元素划分到右侧，条件2天然成立
		if i < m {
			nums1RightMin = nums1[i]
		}

		nums2RightMin := math.MaxInt // j 同理, 注意是j<n
		if j < n {
			nums2RightMin = nums2[j]
		}

		// 判断条件1和2是否满足
		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			// 找到完美分割，根据奇偶给出中位数
			if (m+n)%2 == 1 {
				return float64(max(nums1LeftMax, nums2LeftMax))
			}
			// 否则为偶数, 注意要/2.0
			return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2.0
		} else if nums1LeftMax > nums2RightMin {
			// 太大了往左缩
			right = i - 1
		} else {
			// 太小了往右缩
			left = i + 1
		}

	}
	return 0.0
}
