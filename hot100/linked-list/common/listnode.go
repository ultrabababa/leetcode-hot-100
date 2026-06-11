package common

type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateList 将一个 int 切片快速转换成单链表
// 比如传入 []int{1, 2, 3}，返回 1 -> 2 -> 3 的头节点
func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	dummy := &ListNode{} // 使用虚拟头节点简化操作
	curr := dummy
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}

	return dummy.Next
}

// ListToArray 将单链表转换回 int 切片
// 方便我们在测试用例中直接使用 assert.Equal 进行数组对比
func ListToArray(head *ListNode) []int {
	res := make([]int, 0)
	curr := head
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}
	return res
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
