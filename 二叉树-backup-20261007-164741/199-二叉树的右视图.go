package main

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 199. 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	ans := make([]int, 0)

	for len(q) > 0 {
		levelSize := len(q) // 当前层的节点个数

		// 内层循环遍历当前层的所有节点
		for i := 0; i < levelSize; i++ {
			// 1. 在循环内部弹出节点
			node := q[0]
			q = q[1:]

			// 2. 如果是当前层的最后一个节点，说明从右侧能看到它
			if i == levelSize-1 {
				ans = append(ans, node.Val)
			}

			// 3. 将下一层的节点推入队列（注意大小写 Left 和 Right）
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return ans
}
