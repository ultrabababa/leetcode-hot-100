import unittest
from solutions.hot100.p084_largest_rectangle_in_histogram import Solution

class TestLargestRectangleArea(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_example_1(self):
        # 示例 1
        self.assertEqual(self.sol.largestRectangleArea([2, 1, 5, 6, 2, 3]), 10)

    def test_example_2(self):
        # 示例 2
        self.assertEqual(self.sol.largestRectangleArea([2, 4]), 4)

    def test_single_element(self):
        # 边界用例：单个元素
        self.assertEqual(self.sol.largestRectangleArea([5]), 5)

    def test_all_same(self):
        # 边界用例：所有柱子高度相同
        self.assertEqual(self.sol.largestRectangleArea([3, 3, 3, 3]), 12)

    def test_ascending(self):
        # 边界用例：单调递增
        self.assertEqual(self.sol.largestRectangleArea([1, 2, 3, 4, 5]), 9)

    def test_descending(self):
        # 边界用例：单调递减
        self.assertEqual(self.sol.largestRectangleArea([5, 4, 3, 2, 1]), 9)

if __name__ == '__main__':
    unittest.main()