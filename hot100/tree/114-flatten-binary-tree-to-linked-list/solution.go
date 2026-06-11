package flatten_binary_tree_to_linked_list

import . "leetcode/hot100/tree/common"

func flatten(root *TreeNode) {
	cur := root

	for cur != nil {
		// 如果左子树不为空，需要进行“移花接木”
		if cur.Left != nil {
			// 1. 找左子树的最右节点 (先序遍历中左子树的最后一个节点)
			predecessor := cur.Left
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}

			// 2. 将原来的右子树，接到 predecessor 的右边
			predecessor.Right = cur.Right

			// 3. 将左子树整个移到右边，左边置空
			cur.Right = cur.Left
			cur.Left = nil
		}

		// 4. 继续处理下一个节点
		// 注意：此时的 cur.Right 已经是原来的左子树了（如果发生了移动的话）
		cur = cur.Right
	}
}
