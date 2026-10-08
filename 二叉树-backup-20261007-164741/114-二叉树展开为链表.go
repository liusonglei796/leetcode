package main

// 114. 二叉树展开为链表
// 给你二叉树的根结点 root ，请你要将它展开为一个单链表。
//分治
func flatten(root *TreeNode) {
    var prev *TreeNode

    // 辅助函数：逆先序遍历（右 -> 左 -> 根）
    var postOrder func(node *TreeNode)
    postOrder = func(node *TreeNode) {
        if node == nil {
            return
        }

        // 1. 先递归右子树，再递归左子树（倒着走）
        postOrder(node.Right)
        postOrder(node.Left)

        // 2. 处理当前节点：串联到 prev 上
        node.Right = prev
        node.Left = nil

        // 3. 更新 prev 为当前节点，留给上一个节点接过来
        prev = node
    }

    postOrder(root)
}