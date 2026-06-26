import random
from typing import List


class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
      return self.merge_sort(nums)


    # 基础版快排, 最坏情况会达到 o(n^2)
    def quik_sort(self, nums: List[int]) -> List[int]:
        def qsort(left: int, right: int):
            if left >= right:
                return

            # 选择中间元素作为基准，尽量避免在接近有序的数组中退化
            pivot = nums[left+(right-left)//2]
            i, j = left, right
            while i <= j:
                while nums[i] < pivot: i += 1
                while nums[j] > pivot: j -= 1

                if i <= j:
                    nums[i], nums[j] = nums[j], nums[i]
                    i += 1
                    j -= 1

            qsort(left, j)
            qsort(i, right)

        qsort(0, len(nums)-1)
        return nums


    # 三路划分版快排, 最坏 o(nlogn)
    def three_path_quick_sort(self, nums: List[int]) -> List[int]:
        def qsort3(left: int, right: int):
            if left >= right:
                return

            # 随机选择基准值
            pivot_idx = random.randint(left, right)
            pivot = nums[pivot_idx]

            # lt 指向小于 pivot 的边界，gt 指向大于 pivot 的边界，i 为当前遍历指针
            lt, i, gt = left, left, right

            while i <= gt:
                if nums[i] < pivot:
                    nums[i], nums[lt] = nums[lt], nums[i]
                    i += 1
                    lt += 1
                elif nums[i] > pivot:
                    nums[i], nums[gt] = nums[gt], nums[i]
                    gt -= 1
                else:
                    i += 1

            # 递归处理左右两部分，跳过中间等于 pivot 的部分
            qsort3(left, lt-1)
            qsort3(gt+1, right)

        qsort3(0, len(nums)-1)
        return nums

    # 堆排序
    def heap_sort(self, nums: List[int]) -> List[int]:
        # 内部函数：维护大顶堆的性质（下沉操作）
        def heapify(n: int, i: int):
            """
            :param n: 当前堆的有效大小（参与堆化计算的元素数量）
            :param i: 当前需要进行下沉判断的父节点索引
            """
            largest = i         # 初始化：假设当前节点 i 是最大的
            left = 2 * i + 1
            right = 2 * i + 2

            # 如果左子节点存在，且其值大于当前假设的最大值
            if left < n and nums[left] > nums[largest]:
                largest = left

            if right < n and nums[right] > nums[largest]:
                largest = right

            # 如果最大值不是最初的父节点 i，说明它比子节点小，需要交换
            if largest != i:
                nums[i], nums[largest] = nums[largest], nums[i]

                # 交换后，子节点的值变了，可能会破坏子树的堆结构
                # 因此需要对交换后的子节点继续进行递归下沉操作
                heapify(n, largest)

        n = len(nums)
        if n <= 1:
            return nums

        # 阶段一：建堆 (Build Max-Heap)
        # n // 2 - 1 是最后一个非叶子节点的索引
        # 我们从最后一个非叶子节点开始，逆序遍历到根节点（索引 0）
        for i in range(n // 2 - 1, -1, -1):
            heapify(n, i)

        # 阶段二：排序 (Sort / Extract Max)
        # 从数组最后一个元素开始，逆序遍历到索引 1
        for i in range(n - 1, 0, -1):
            # 将当前堆顶（最大值 nums[0]）交换到数组末尾 nums[i]
            nums[0], nums[i] = nums[i], nums[0]
            # 交换后，排好序的元素被排除在堆的有效大小之外 (有效大小变为 i)
            # 此时新的堆顶 nums[0] 破坏了堆结构，需要从上到下重新堆化一次
            heapify(i, 0)

        return nums



    # 归并排序
    def merge_sort(self, nums: list[int]) -> List[int]:
        if len(nums) <= 1:
            return nums

        mid = len(nums) // 2

        left = self.merge_sort(nums[:mid])
        right = self.merge_sort(nums[mid:])

        return self._merge(left, right)


    def _merge(self, left: List[int], right: List[int]) -> List[int]:
        res = []
        i = 0
        j = 0

        # 双指针遍历，比较两个数组的当前元素
        while i < len(left) and j < len(right):
            # 注意这里的 <=，等号保证了算法的“稳定性”
            # 即当元素相等时，优先保留左边（原序列靠前）的元素
            if left[i] <= right[j]:
                res.append(left[i])
                i += 1
            else:
                res.append(right[j])
                j += 1

        # 当其中一个数组被遍历完后，另一个数组可能还有剩余元素
        # 因为输入的数组本身就是有序的，直接将剩余部分追加到尾部即可
        if i < len(left):
            res.extend(left[i:])
        if j < len(right):
            res.extend(right[j:])

        return res