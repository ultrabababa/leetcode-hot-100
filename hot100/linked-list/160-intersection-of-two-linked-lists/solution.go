package intersection_of_two_linked_lists

import . "leetcode/hot100/linked-list/common"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 如果有任何一个链表为空，直接返回 nil
	if headA == nil || headB == nil {
		return nil
	}

	pA := headA
	pB := headB

	// 当 pA 和 pB 指向同一个节点时（相遇）停止循环
	// 即使它们都不相交，最后也会同时变成 nil，nil == nil 成立，跳出循环
	for pA != pB {
		// pA 走完了自己的路，去走 B 的路
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		// pB 走完了自己的路，去走 A 的路
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	// 此时 pA 就是相交的节点（或者 nil）
	return pA
}
