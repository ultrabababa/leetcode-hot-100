package intersection_of_two_linked_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

import . "leetcode/hot100/linked-list/common"

func TestGetIntersectionNode(t *testing.T) {

	// 辅助函数：获取链表尾部节点
	getTail := func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		curr := head
		for curr.Next != nil {
			curr = curr.Next
		}
		return curr
	}

	t.Run("Example 1: Intersected at 8", func(t *testing.T) {
		headA := CreateList([]int{4, 1})
		headB := CreateList([]int{5, 6, 1})
		common := CreateList([]int{8, 4, 5})

		// 将 common 拼接到 A 和 B 的尾部，形成真实的物理相交
		getTail(headA).Next = common
		getTail(headB).Next = common

		res := getIntersectionNode(headA, headB)
		assert.Equal(t, common, res)
	})

	t.Run("Example 2: Intersected at 2", func(t *testing.T) {
		headA := CreateList([]int{1, 9, 1})
		headB := CreateList([]int{3})
		common := CreateList([]int{2, 4})

		getTail(headA).Next = common
		getTail(headB).Next = common

		res := getIntersectionNode(headA, headB)
		assert.Equal(t, common, res)
	})

	t.Run("Example 3: No intersection", func(t *testing.T) {
		headA := CreateList([]int{2, 6, 4})
		headB := CreateList([]int{1, 5})

		res := getIntersectionNode(headA, headB)
		assert.Nil(t, res) // 期望返回 nil
	})
}
