package merge_two_sorted_lists

import . "leetcode/hot100/linked-list/common"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 1. 创建哨兵节点 (Dummy Head)
	// 它的作用是避免处理头节点为空的边界情况
	dummy := &ListNode{Val: -1}

	// cur 指针用于追踪新链表的尾部
	cur := dummy

	// 2. 当两个链表都有节点时，进行比较和拼接
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1   // 接上 list1
			list1 = list1.Next // list1 指针后移
		} else {
			cur.Next = list2   // 接上 list2
			list2 = list2.Next // list2 指针后移
		}
		cur = cur.Next // cur 指针永远指向最后那个节点
	}

	// 3. 处理剩余部分
	// 循环结束时，必定有一个链表还没遍历完（或者两个都刚好完了）
	// 直接把剩下的接在后面就行
	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}

	// 4. 返回哨兵节点的下一个节点，即真正的头节点
	return dummy.Next
}
