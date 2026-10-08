package main

import "math"

// 124. 二叉树中的最大路径和
// 给你一个二叉树的根节点 root ，返回其 最大路径和。
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 初始化为极小值，防止整棵树都是负数时答案被 0 压住
	ans := math.MinInt

	// 关键：dfs 有两个职责
	//   1. 返回以当前节点为起点、向下的单边最大路径和（只能选左或右一边）
	//   2. 顺带更新 ans —— 记录"拐点"在当前节点的全局最优解
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 子树贡献为负就舍弃，与 0 取最大
		leftGain := max(dfs(node.Left), 0)
		rightGain := max(dfs(node.Right), 0)

		// 拐点在当前节点：左右都能走，node.Val 也不能丢
		ans = max(ans, node.Val+leftGain+rightGain)

		// 汇报给父节点时只能选一条边继续往上延伸
		return node.Val + max(leftGain, rightGain)
	}

	dfs(root)
	return ans
}
