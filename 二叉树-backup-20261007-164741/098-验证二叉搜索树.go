package main


import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    isValid := true

    var dfs func(node *TreeNode, low, max int)
    dfs = func(node *TreeNode, low, max int) {
        // 如果已经发现不合法，或者节点为空，直接返回，不需要返回具体值
        if !isValid || node == nil {
            return
        }
        // 发现违规，打上标记
        if node.Val <= low || node.Val >= max {
            isValid = false
            return
        }
        // 继续向下传递约束
        dfs(node.Left, low, node.Val)
        dfs(node.Right, node.Val, max)
    }

    dfs(root, math.MinInt, math.MaxInt)
    return isValid
}