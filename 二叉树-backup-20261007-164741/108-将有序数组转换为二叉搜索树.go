package main

// 108. 将有序数组转换为二叉搜索树
// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)
		return root
	}
	return dfs(0, len(nums)-1)
}
