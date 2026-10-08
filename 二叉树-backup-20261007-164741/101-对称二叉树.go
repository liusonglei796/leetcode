package main

// 101. 对称二叉树
// 给你一个二叉树的根节点 root，检查它是否轴对称。
// 自顶向下模型，需要辅助函数
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(p, q *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}

	return dfs(root.Left, root.Right)
}
