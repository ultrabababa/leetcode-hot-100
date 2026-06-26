import random
from typing import List


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        def quick_select(arr: List[int], k: int) -> int:
            # 1. 随机选择基准值
            pivot = random.choice(arr)

            # 2. 三路划分：大于、小于、等于基准值
            big = [x for x in arr if x > pivot]
            small = [x for x in arr if x < pivot]
            equal = [x for x in arr if x == pivot]

            # 3. 判断第 k 大元素在哪个区间
            if k <= len(big):
                return quick_select(big, k)
            elif k > (len(big) + len(equal)):
                return quick_select(small, k - len(big) - len(equal))
            else:
                return pivot

        return quick_select(nums, k)
