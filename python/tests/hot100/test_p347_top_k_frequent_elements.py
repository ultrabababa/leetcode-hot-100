import unittest
from solutions.hot100.p347_top_k_frequent_elements import Solution

class TestTopKFrequent(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_1(self):
        nums = [1, 1, 1, 2, 2, 3]
        k = 2
        # 注意：题目说明可以按任意顺序返回，因此推荐使用 set 进行无序断言
        self.assertEqual(set(self.solution.topKFrequent(nums, k)), {1, 2})

    def test_example_2(self):
        nums = [1]
        k = 1
        self.assertEqual(self.solution.topKFrequent(nums, k), [1])

    def test_example_3(self):
        nums = [1, 2, 1, 2, 1, 2, 3, 1, 3, 2]
        k = 2
        self.assertEqual(set(self.solution.topKFrequent(nums, k)), {1, 2})

    def test_negative_numbers(self):
        nums = [-1, -1, -2, -2, -2, -3]
        k = 1
        self.assertEqual(self.solution.topKFrequent(nums, k), [-2])

if __name__ == '__main__':
    unittest.main()