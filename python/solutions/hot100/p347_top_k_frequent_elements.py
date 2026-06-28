import collections
from typing import List


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        # 1. 使用哈希表统计每个元素的出现频率
        # 例如 nums = [1,1,1,2,2,3], freq_map = {1: 3, 2: 2, 3: 1}
        freq_map = collections.Counter(nums)

        # 2. 创建桶数组，索引代表频率，由于频率最高不可能超过数组长度，所以长度为 len(nums) + 1
        # buckets[i] 存储出现频率正好为 i 的所有元素
        buckets = [[] for _ in range(len(nums)+1)]
        for num, freq in freq_map.items():
            buckets[freq].append(num)

        # 3. 倒序遍历桶数组，收集频率最高的前 k 个元素
        result = []
        for i in range(len(buckets)-1, 0, -1):
            if buckets[i]:
                # 如果当前频率桶中有元素，将其加入结果集
                result.extend(buckets[i])
                # 如果结果集已满 k 个，直接返回
                if len(result) >= k:
                    return result[:k]

        return result