package merge_k_sorted_lists

import (
	"container/heap"
	. "leetcode/hot100/linked-list/common"
)

// 解法一
// mergeKLists 主函数
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// 1. 初始化最小堆
	h := &ListNodeHeap{}
	heap.Init(h)

	// 2. 将所有链表的【头节点】放入堆中
	// 注意：只放头节点，不需要把整个链表放进去
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	// 3. 建立虚拟头节点
	dummy := &ListNode{Val: -1}
	cur := dummy

	// 4. 不断从堆中取出最小节点，接在结果链表后面
	for h.Len() > 0 {
		// Pop 出当前最小的节点
		minNode := heap.Pop(h).(*ListNode)

		cur.Next = minNode // 接上
		cur = cur.Next     // 指针后移

		// 关键一步：如果这个节点还有下一个节点，把它推入堆中
		// 相当于“这个班级的下一个人补位”
		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
	}

	return dummy.Next
}

// ==========================================
// 下面是 Go 语言标准库 container/heap 的模板代码
// 在面试中，这部分代码通常需要手写
// ==========================================

// 定义一个类型，实现 heap.Interface 接口
type ListNodeHeap []*ListNode

func (h *ListNodeHeap) Len() int { return len(*h) }

// Less 决定是大顶堆还是小顶堆。这里用 < 号，所以是小顶堆
func (h *ListNodeHeap) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }

func (h *ListNodeHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

// Push 和 Pop 必须使用指针接收者，因为切片长度会变
func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 解法2
// 这种写法也是 O(N log K)，而且不用写堆的模板
func mergeKListsByDivideAndConquer(lists []*ListNode) *ListNode {
	amount := len(lists)
	if amount == 0 {
		return nil
	}
	if amount == 1 {
		return lists[0]
	}

	mid := amount / 2
	left := mergeKListsByDivideAndConquer(lists[:mid])
	right := mergeKListsByDivideAndConquer(lists[mid:])

	// 复用之前的 mergeTwoLists 函数
	return mergeTwoLists(left, right)
}

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
