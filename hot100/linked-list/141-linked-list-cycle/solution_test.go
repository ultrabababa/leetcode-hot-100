package linked_list_cycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 引入公共包
	. "leetcode/hot100/linked-list/common"
)

func TestHasCycle(t *testing.T) {
	// 专门为这道题写的辅助函数：创建带环链表
	// pos 表示尾节点指向的索引，-1 表示无环
	createCycleList := func(nums []int, pos int) *ListNode {
		head := CreateList(nums)
		if pos == -1 {
			return head
		}

		// 1. 找到需要连接的目标节点 (cycleNode)
		cycleNode := head
		for i := 0; i < pos; i++ {
			cycleNode = cycleNode.Next
		}

		// 2. 找到尾节点 (tail)
		tail := head
		for tail.Next != nil {
			tail = tail.Next
		}

		// 3. 制造环：尾巴咬住目标节点
		tail.Next = cycleNode

		return head
	}

	type testCase struct {
		name     string
		head     []int
		pos      int
		expected bool
	}

	tests := []testCase{
		{
			name:     "Example 1",
			head:     []int{3, 2, 0, -4},
			pos:      1, // 尾部接到 index 1 (值2)
			expected: true,
		},
		{
			name:     "Example 2",
			head:     []int{1, 2},
			pos:      0, // 尾部接到 index 0 (值1)
			expected: true,
		},
		{
			name:     "Example 3 - No Cycle",
			head:     []int{1},
			pos:      -1,
			expected: false,
		},
		{
			name:     "Empty List",
			head:     []int{},
			pos:      -1,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			listHead := createCycleList(tc.head, tc.pos)
			result := hasCycle(listHead)
			assert.Equal(t, tc.expected, result)
		})
	}
}
