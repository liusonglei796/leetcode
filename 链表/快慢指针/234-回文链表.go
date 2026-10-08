package lc

// 234. 回文链表（简单）
//
// 给你一个单链表 head，请你判断链表是否是回文链表。
// 如果是，返回 true；否则，返回 false。
// 进阶：能否用 O(1) 时间复杂度解决？
//
// 示例 1：head = [1,2,2,1]，输出 true
// 示例 2：head = [1,2]，输出 false
//
// 思路：快慢指针找中点 → 反转后半段 → 与前半段逐个比较 → 还原链表。
// 用到的反转函数见本文件下方的 reverseList。

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	right := slow.Next
	slow.Next = nil
	r := reverseList(right)

	p1, p2 := head, r
	for p2 != nil {
		if p1.Val != p2.Val {
			slow.Next = reverseList(r)
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	slow.Next = reverseList(r)
	return true
}

// reverseList 迭代反转单链表（辅助函数）

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
