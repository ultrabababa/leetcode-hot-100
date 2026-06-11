package linked_list_cycle_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 使用点导入，利用我们之前的 common 包
	. "leetcode/hot100/linked-list/common"
)

func TestDetectCycle(t *testing.T) {
	// 辅助函数：创建带环链表，并返回【环的入口节点】
	// 如果 pos = -1，返回 nil
	createCycleListAndGetEntry := func(nums []int, pos int) (*ListNode, *ListNode) {
		head := CreateList(nums)
		if pos == -1 {
			return head, nil
		}

		// 1. 找到入口节点
		entryNode := head
		for i := 0; i < pos; i++ {
			entryNode = entryNode.Next
		}

		// 2. 找到尾节点
		tail := head
		for tail.Next != nil {
			tail = tail.Next
		}

		// 3. 制造环
		tail.Next = entryNode

		return head, entryNode
	}

	type testCase struct {
		name string
		nums []int
		pos  int
	}

	tests := []testCase{
		{
			name: "Example 1",
			nums: []int{3, 2, 0, -4},
			pos:  1, // 环入口是 index 1 (val 2)
		},
		{
			name: "Example 2",
			nums: []int{1, 2},
			pos:  0, // 环入口是 index 0 (val 1)
		},
		{
			name: "Example 3 - No Cycle",
			nums: []int{1},
			pos:  -1,
		},
		{
			name: "No Cycle Long List",
			nums: []int{1, 2, 3, 4},
			pos:  -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, expectedEntry := createCycleListAndGetEntry(tc.nums, tc.pos)

			// 运行算法
			result := detectCycle(head)

			// 断言：我们要比较的是【节点对象本身】，即内存地址是否相同
			assert.Equal(t, expectedEntry, result)
		})
	}
}
