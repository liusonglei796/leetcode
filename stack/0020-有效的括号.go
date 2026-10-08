package stack

// 20. 有效的括号
// 简单

// 给定一个只包含 '(', ')', '{', '}', '[' 和 ']' 的字符串 s，
// 判断字符串是否有效。

// 有效字符串需满足：
// 1. 左括号必须用相同类型的右括号闭合。
// 2. 左括号必须以正确的顺序闭合。
// 3. 每个右括号都有一个对应的相同类型的左括号。

// 示例 1:
// 输入: s = "()"
// 输出: true

// 示例 2:
// 输入: s = "()[]{}"
// 输出: true

// 示例 3:
// 输入: s = "(]"
// 输出: false

// 提示:
// 1 <= s.length <= 10^4
// s 仅由括号 '()[]{}' 组成


func isValid(s string) bool {
	m := map[byte]byte{
		']': '[',
		')': '(',
		'}': '{',
	}
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			// 左括号，压栈
			stack = append(stack, s[i])
		} else {
			// 右括号，检查栈顶
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1] // 弹出栈顶
		}
	}

	return len(stack) == 0
}
