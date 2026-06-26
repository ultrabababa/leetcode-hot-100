from typing import List


class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        # 引入哨兵，首尾各加一个 0
        new_heights = [0] + heights + [0]

        # 栈里面存放索引，维护高度单调递增
        stack = []
        max_area = 0

        for i, h in enumerate(new_heights):
            # 当破坏了单调递增规则，说明栈顶元素找到了右边界，可以开始结算
            while stack and new_heights[stack[-1]] > h:
                # 1. 弹出栈顶，以此高度作为当前计算矩形的瓶颈高度
                cur_height_idx = stack.pop()
                cur_height = new_heights[cur_height_idx]

                # 2. 此时的新栈顶即为左边界，当前遍历的索引 i 即为右边界
                left_bound = stack[-1]
                right_bound = i
                cur_width = right_bound - left_bound - 1

                # 3. 计算并更新最大面积
                max_area = max(max_area, cur_height * cur_width)

            # 满足单调递增，当前索引入栈
            stack.append(i)

        return max_area