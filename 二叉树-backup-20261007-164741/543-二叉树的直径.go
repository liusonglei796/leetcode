package main

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	// dfs 返回以当前节点为起点的单侧路径上的「节点总数」（即单侧深度）
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 1. 拿到左右子树各自的单侧节点数
		leftNodes := dfs(node.Left)
		rightNodes := dfs(node.Right)

		// 2. 计算穿过当前拐点节点的「整条路径总节点数」：
		// 左子树节点数 + 右子树节点数 + 当前节点自身 (1 个)
		totalNodes := leftNodes + rightNodes + 1

		// 3. 树路径的物理边数 = 节点数 - 1
		currentEdges := totalNodes - 1

		// 4. 更新全局最大直径（边数）
		if currentEdges > maxDiameter {
			maxDiameter = currentEdges
		}

		// 5. 向父节点汇报当前节点向下的单侧最大节点数
		return max(leftNodes, rightNodes) + 1
	}

	dfs(root)
	return maxDiameter
}