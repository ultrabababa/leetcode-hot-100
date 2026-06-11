package linked_list_cycle

import . "leetcode/hot100/linked-list/common"

func hasCycle(head *ListNode) bool {
	// 如果链表为空或只有一个节点，肯定没环
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	// 只要快指针还没跑到终点
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢走一步
		fast = fast.Next.Next // 快走两步

		// 撞上了！说明有环
		if slow == fast {
			return true
		}
	}

	// 快指针跑到了 nil，说明是直路，没环
	return false
}
