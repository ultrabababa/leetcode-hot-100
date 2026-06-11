package swap_nodes_in_pairs

import . "leetcode/hot100/linked-list/common"

func swapPairs(head *ListNode) *ListNode {
	// 1. 无脑上 Dummy Node
	dummy := &ListNode{Next: head}
	cur := dummy

	// 2. 循环条件：必须保证后面还有【两个】节点可以交换
	// 如果只剩一个节点（奇数个），或者没节点了，就结束
	for cur.Next != nil && cur.Next.Next != nil {
		// 提前保存要交换的两个节点，防止后面指针乱飞找不到了
		node1 := cur.Next
		node2 := cur.Next.Next

		// 开始交换（画图最好懂）
		cur.Next = node2        // 步骤 1：cur 指向第二个节点
		node1.Next = node2.Next // 步骤 2：第一个节点指向第三个节点（防止断链）
		node2.Next = node1      // 步骤 3：第二个节点反过来指向第一个节点

		// 移动 cur 指针，准备下一对
		// 此时 node1 已经成了这对里的最后一个节点，所以 cur 移到 node1
		cur = node1
	}

	// 3. 返回真实的新头节点
	return dummy.Next
}
