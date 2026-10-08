package main

// 236. 二叉树的最近公共祖先
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
   if root==nil || root==p ||root==q{
	return root
   }
   left:=lowestCommonAncestor(root.Left,p,q)
   right:=lowestCommonAncestor(root.Right,p,q)
   if left!=nil&&right!=nil{
	return root
   }
   if left!=nil{
	return left
   }
   return right
}
