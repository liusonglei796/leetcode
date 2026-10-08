package main

// 230. 二叉搜索树中第K小的元素
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
func kthSmallest(root *TreeNode, k int) int {
	ans := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || k <= 0 {
			return
		}
		dfs(node.Left)
		k--
		if k == 0 {
			ans = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
