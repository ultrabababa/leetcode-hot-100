package copy_list_with_random_pointer

import (
	"testing"

	. "leetcode/hot100/linked-list/common"

	"github.com/stretchr/testify/assert"
)

func TestCopyRandomList(t *testing.T) {
	// 辅助函数：根据 [[val, random_index], ...] 构造链表
	buildList := func(data [][]interface{}) *Node {
		if len(data) == 0 {
			return nil
		}

		// 1. 先创建所有节点，存入数组方便索引
		nodes := make([]*Node, len(data))
		for i, item := range data {
			val := item[0].(int)
			nodes[i] = &Node{Val: val}
		}

		// 2. 再次遍历，连接 Next 和 Random
		for i, item := range data {
			// 连接 Next
			if i < len(data)-1 {
				nodes[i].Next = nodes[i+1]
			}
			// 连接 Random
			randomIndex := item[1]
			if randomIndex != nil {
				idx := randomIndex.(int)
				nodes[i].Random = nodes[idx]
			}
		}
		return nodes[0]
	}

	// 辅助函数：验证两个链表是否完全一样（深拷贝）
	// 返回 true 表示值和结构相同，但内存地址不同
	checkDeepCopy := func(headA, headB *Node) bool {
		iterA, iterB := headA, headB
		// 简单的 map 来检查 random 指向的一致性
		mapA := make(map[*Node]int) // 记录 A 中节点的 index
		idx := 0
		for p := headA; p != nil; p = p.Next {
			mapA[p] = idx
			idx++
		}

		for iterA != nil && iterB != nil {
			// 1. 值必须相同
			if iterA.Val != iterB.Val {
				return false
			}
			// 2. 内存地址必须不同（深拷贝核心）
			if iterA == iterB {
				return false
			}

			// 3. 检查 Random 指向的相对位置是否一致
			if iterA.Random != nil {
				if iterB.Random == nil {
					return false
				}
				// A 的 Random 指向第 k 个节点，B 的 Random 也必须指向第 k 个
				idxA := mapA[iterA.Random]

				// 找到 B 的 Random 对应的 index (这里简化处理，手动遍历一下B找位置)
				idxB := 0
				found := false
				for p := headB; p != nil; p = p.Next {
					if p == iterB.Random {
						found = true
						break
					}
					idxB++
				}
				if !found || idxA != idxB {
					return false
				}
			} else if iterB.Random != nil {
				return false
			}

			iterA = iterA.Next
			iterB = iterB.Next
		}
		return iterA == nil && iterB == nil
	}

	type testCase struct {
		name string
		data [][]interface{} // 使用 interface{} 兼容 int 和 nil
	}

	tests := []testCase{
		{
			name: "Example 1",
			data: [][]interface{}{
				{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0},
			},
		},
		{
			name: "Example 2",
			data: [][]interface{}{
				{1, 1}, {2, 1},
			},
		},
		{
			name: "Example 3",
			data: [][]interface{}{
				{3, nil}, {3, 0}, {3, nil},
			},
		},
		{
			name: "Empty List",
			data: [][]interface{}{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := buildList(tc.data)
			newHead := copyRandomList(head)

			if len(tc.data) == 0 {
				assert.Nil(t, newHead)
			} else {
				assert.True(t, checkDeepCopy(head, newHead))
			}
		})
	}
}
