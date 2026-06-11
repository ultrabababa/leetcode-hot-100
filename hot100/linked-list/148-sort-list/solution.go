package sort_list

import . "leetcode/hot100/linked-list/common"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := splitLists(head)
	left := sortList(head)
	right := sortList(mid)

	return mergeLists(left, right)
}

func splitLists(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil
	return mid
}

func mergeLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
