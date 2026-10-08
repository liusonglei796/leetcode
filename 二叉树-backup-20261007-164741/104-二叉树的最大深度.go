package main

// 104. 二叉树的最大深度
// 给定一个二叉树 root ，返回其最大深度。
//分治
func maxDepth(root *TreeNode) int {
     if root==nil{
		return 0
	 }  
	left:= maxDepth(root.Left)
	Right:=maxDepth(root.Right)
    return max(left,Right)+1
}
