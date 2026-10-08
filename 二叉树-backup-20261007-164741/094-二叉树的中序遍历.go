package main

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ===== 94. 二叉树的中序遍历：三种切片传递写法 =====
// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。

// 1. 传切片指针 + 解引用
// 特点：适合多级嵌套函数调用，或显式将逻辑封装在独立的外部函数中。
func inorderTraversal1(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *TreeNode, res *[]int)
	dfs = func(node *TreeNode, res *[]int) {
		if node == nil {
			return
		}
		dfs(node.Left, res)
		*res = append(*res, node.Val) // append 的第一个参数必须是切片本身，不是 *[]int
		dfs(node.Right, res)
	}
	dfs(root, &ans)
	return ans
}

// 2. 利用闭包隐式传递（最推荐的 Go 风格）
// 特点：没有切片形参干扰，永远修改外层唯一的 ans 变量，避开扩容"分家"问题。
func inorderTraversal2(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// 3. 传普通切片 + 返回新切片（函数式风格）
// 特点：遵循"谁修改，谁返回"，扩容后新指针通过 return 层层传回，不丢数据。
func inorderTraversal3(root *TreeNode) []int {
	return dfs3(root, []int{})
}

func dfs3(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = dfs3(node.Left, res)
	res = append(res, node.Val)
	res = dfs3(node.Right, res)
	return res
}
