import unittest
from solutions.hot100.p912_sort_an_array import Solution


class TestSortArray(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_example_1(self):
        # 示例 1：常规无序数组
        self.assertEqual(self.sol.sortArray([5, 2, 3, 1]), [1, 2, 3, 5])

    def test_example_2(self):
        # 示例 2：包含重复元素和 0
        self.assertEqual(self.sol.sortArray([5, 1, 1, 2, 0, 0]), [0, 0, 1, 1, 2, 5])

    def test_already_sorted(self):
        # 边界用例：已经升序
        self.assertEqual(self.sol.sortArray([1, 2, 3, 4, 5]), [1, 2, 3, 4, 5])

    def test_reverse_sorted(self):
        # 边界用例：完全降序
        self.assertEqual(self.sol.sortArray([5, 4, 3, 2, 1]), [1, 2, 3, 4, 5])

    def test_negative_numbers(self):
        # 边界用例：包含负数
        self.assertEqual(self.sol.sortArray([-3, -1, -5, 2, 0]), [-5, -3, -1, 0, 2])


if __name__ == '__main__':
    unittest.main()