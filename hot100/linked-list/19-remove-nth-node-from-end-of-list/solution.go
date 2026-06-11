package remove_nth_node_from_end_of_list

import . "leetcode/hot100/linked-list/common"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1. 设置虚拟头节点，这能完美解决“删除头节点”的边界问题
	dummy := &ListNode{Next: head}
	fast := dummy
	slow := dummy

	// 2. fast 先走 n + 1 步
	// 为什么要走 n+1 步？因为只有这样，最后 slow 才会停在目标节点的前一个位置
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// 3. fast 和 slow 同步前进，直到 fast 变成 nil (也就是跑出了链表尾部)
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 4. 此时 slow 刚好指在【倒数第 N 个节点的前驱】
	// 执行删除操作：越过被删节点，直接连接下一个
	slow.Next = slow.Next.Next

	// 5. 返回真正的头节点
	return dummy.Next
}
