package copy_list_with_random_pointer

import . "leetcode/hot100/linked-list/common"

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 1. 创建哈希表：key 是原节点，value 是新节点
	nodeMap := make(map[*Node]*Node)

	// 2. 第一轮遍历：只创建新节点，存入 map
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val}
		nodeMap[cur] = newNode
		cur = cur.Next
	}
	// nodeMap[nil] = nil // 甚至可以不加，当访问 map 中不存在的 key 时，会返回该 value 类型的零值

	for p := head; p != nil; p = p.Next {
		curNode := nodeMap[p]
		// 即使 p.Next 是 nil，nodeMap[nil] 也会返回 nil
		// 即使 p.Random 是 nil，nodeMap[nil] 也会返回 nil
		curNode.Next = nodeMap[p.Next]
		curNode.Random = nodeMap[p.Random]
	}
	return nodeMap[head]
}
