import unittest
from solutions.hot100.p295_find_median_from_data_stream import MedianFinder

class TestMedianFinder(unittest.TestCase):
    def test_example_1(self):
        # 对应题目示例 1
        median_finder = MedianFinder()
        median_finder.addNum(1)
        median_finder.addNum(2)
        self.assertEqual(median_finder.findMedian(), 1.5)
        median_finder.addNum(3)
        self.assertEqual(median_finder.findMedian(), 2.0)

    def test_negative_numbers(self):
        # 测试包含负数的情况
        median_finder = MedianFinder()
        median_finder.addNum(-1)
        self.assertEqual(median_finder.findMedian(), -1.0)
        median_finder.addNum(-2)
        self.assertEqual(median_finder.findMedian(), -1.5)
        median_finder.addNum(-3)
        self.assertEqual(median_finder.findMedian(), -2.0)
        median_finder.addNum(-4)
        self.assertEqual(median_finder.findMedian(), -2.5)
        median_finder.addNum(5)
        self.assertEqual(median_finder.findMedian(), -2.0)

    def test_duplicates(self):
        # 测试重复数字
        median_finder = MedianFinder()
        median_finder.addNum(2)
        median_finder.addNum(2)
        self.assertEqual(median_finder.findMedian(), 2.0)
        median_finder.addNum(2)
        self.assertEqual(median_finder.findMedian(), 2.0)

if __name__ == '__main__':
    unittest.main()