package palindrome_linked_list

import . "leetcode/hot100/linked-list/common"

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 1. 快慢指针找中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 此时 slow 指向了后半部分的起点 (如果节点数是奇数，slow在正中间)

	// 2. 反转后半部分链表 (原地反转)
	var pre *ListNode
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// 反转结束后，pre 成为了后半段反转后的新头节点

	// 3. 比较前后两部分
	p1 := head // 前半部分起点
	p2 := pre  // 后半部分逆序后的起点
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}
