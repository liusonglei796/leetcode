package stack

// 394. 字符串解码
// 中等

// 给定一个经过编码的字符串，返回它解码后的字符串。

// 编码规则为: k[encoded_string]，表示其中的 encoded_string 正好重复 k 次。
// 注意 k 保证为正整数。

// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，
// 且输入的方括号总是符合格式要求的。

// 此外，你可以认为原始数据不包含数字，所有数字只表示重复次数 k，
// 例如不会出现像 3a 或 2[4] 的输入。

// 示例 1:
// 输入: s = "3[a2[c]]"
// 输出: "accaccacc"

// 示例 2:
// 输入: s = "2[abc]3[cd]ef"
// 输出: "abcabccdcdcdef"

// 示例 3:
// 输入: s = "abc3[cd]xyz"
// 输出: "abccdcdcdxyz"

// 提示:
// 1 <= s.length <= 30
// s 由小写英文字母、数字和方括号 '[]' 组成
// s 保证是一个 有效的 输入。
// s 中所有整数的范围为 [1, 300]

func decodeString(s string) string {
	var numStack []int       // 存数字
	var strStack []string    // 存字符串
	currStr := ""
	currNum := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			// 解析数字
			currNum = currNum*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			// 遇到 '['，压栈，重置当前状态
			numStack = append(numStack, currNum)
			strStack = append(strStack, currStr)
			currNum = 0
			currStr = ""
		} else if s[i] == ']' {
			// 遇到 ']'，弹出栈，拼接字符串
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			currStr = prevStr + repeat(currStr, num)
		} else {
			// 遇到字母，追加到当前字符串
			currStr += string(s[i])
		}
	}

	return currStr
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
