package reverse_linked_list

import . "leetcode/hot100/linked-list/common"

// last review 6.17

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		nextTemp := curr.Next // 1. 先把下一个节点存起来，防止失联
		curr.Next = prev      // 2. 核心：把当前节点的箭头反转，指向前一个节点
		prev = curr           // 3. prev 指针往前挪一步
		curr = nextTemp       // 4. curr 指针往前挪一步
	}

	// 当 curr 为 nil 时，prev 刚好停在原链表的最后一个节点，也就是新链表的头节点
	return prev
}
