package reverse_nodes_in_k_group

import . "leetcode/hot100/linked-list/common"

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy // pre 永远指向当前准备翻转的子链表的前驱节点

	for head != nil {
		// 1. 探路：找到当前组的尾部节点 (end)
		end := pre
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				// 如果剩余节点不足 k 个，直接返回，保持原样
				return dummy.Next
			}
		}

		// 2. 截断：记录下一组的开头，并把当前组独立出来
		nextGroup := end.Next
		start := pre.Next      // 当前组的开头
		end.Next = nil         // 切断链表，制造一个独立的子链表

		// 3. 翻转：翻转这个独立的子链表
		// 翻转后，end 会变成头，start 会变成尾
		pre.Next = reverseList(start)

		// 4. 缝合：把翻转后的子链表和后面的大部队接上
		start.Next = nextGroup

		// 5. 推进：更新 pre 和 head，准备下一组
		pre = start            // start 现在是翻转后的尾巴，它将作为下一组的 pre
		head = nextGroup       // head 移到下一组的开头
	}

	return dummy.Next
}

// 经典的翻转链表辅助函数 (和 第 206 题一模一样)
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre // 返回翻转后的新头节点
}
