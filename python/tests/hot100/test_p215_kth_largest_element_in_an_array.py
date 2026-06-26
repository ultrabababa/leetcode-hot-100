import unittest
from solutions.hot100.p215_kth_largest_element_in_an_array import Solution

class TestFindKthLargest(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_example_1(self):
        # 示例 1
        self.assertEqual(self.sol.findKthLargest([3, 2, 1, 5, 6, 4], 2), 5)

    def test_example_2(self):
        # 示例 2：包含重复元素
        self.assertEqual(self.sol.findKthLargest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4), 4)

    def test_all_same_elements(self):
        # 边界用例：所有元素相同
        self.assertEqual(self.sol.findKthLargest([2, 2, 2, 2, 2], 3), 2)

    def test_sorted_ascending(self):
        # 边界用例：已经升序排列
        self.assertEqual(self.sol.findKthLargest([1, 2, 3, 4, 5, 6], 1), 6)

    def test_sorted_descending(self):
        # 边界用例：已经降序排列
        self.assertEqual(self.sol.findKthLargest([6, 5, 4, 3, 2, 1], 6), 1)

if __name__ == '__main__':
    unittest.main()