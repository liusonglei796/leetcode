package main

// 102. 二叉树的层序遍历
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历。
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    ans := make([][]int, 0)
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)          // 1. 核心：必须先固定当前层的节点数量
        levelValues := make([]int, 0, levelSize)

        for i := 0; i < levelSize; i++ { // 2. 严格循环当前层的次数
            node := queue[0]
            queue = queue[1:]

            levelValues = append(levelValues, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        ans = append(ans, levelValues)   // 3. 将整层结果存入 ans
    }

    return ans
}