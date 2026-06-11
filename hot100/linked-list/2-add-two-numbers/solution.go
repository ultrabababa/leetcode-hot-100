package add_two_numbers

import . "leetcode/hot100/linked-list/common"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 标准起手式：Dummy Node
	dummy := &ListNode{Val: -1}
	cur := dummy

	carry := 0 // 进位

	// 2. 只要 l1 有值，或者 l2 有值，或者还有进位没处理，就继续
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry // 先加上进位

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// 计算新节点的值和新的进位
		carry = sum / 10    // 比如 15 / 10 = 1
		nodeVal := sum % 10 // 比如 15 % 10 = 5

		// 接在新链表后面
		cur.Next = &ListNode{Val: nodeVal}
		cur = cur.Next
	}

	return dummy.Next
}
