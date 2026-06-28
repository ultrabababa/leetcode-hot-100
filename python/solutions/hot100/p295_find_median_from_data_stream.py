import heapq


class MedianFinder:

    def __init__(self):
        # 小顶堆，存储较大的一半数据
        self.min_heap = []
        # 大顶堆，存储较小的一半数据（Python中存负数来模拟大顶堆）
        self.max_heap = []

    def addNum(self, num: int) -> None:
        # 1. 默认先压入大顶堆（存负值）
        heapq.heappush(self.max_heap, -num)
        # 2. 将大顶堆的最大值（取反后为最小的负数）弹出，压入小顶堆，确保数据大小关系正确
        max_in_small_part = -heapq.heappop(self.max_heap)
        heapq.heappush(self.min_heap, max_in_small_part)

        # 3. 维护堆的数量平衡，保证大顶堆数量 >= 小顶堆数量
        if len(self.min_heap) > len(self.max_heap):
            min_in_large_part = heapq.heappop(self.min_heap)
            heapq.heappush(self.max_heap, -min_in_large_part)


    def findMedian(self) -> float:
        # 如果总长度是奇数，大顶堆元素必定比小顶堆多 1，返回大顶堆堆顶（记得取反）
        if len(self.max_heap) > len(self.min_heap):
            return float(-self.max_heap[0])

        return (-self.max_heap[0] + self.min_heap[0]) / 2.0

# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()