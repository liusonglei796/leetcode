package lc

// 148. 排序链表（中等）
//
// 给你链表的头节点 head，请将链表按升序排列并返回排序后的链表。
//
// 示例 1：head = [4,2,1,3]，输出 [1,2,3,4]
// 示例 2：head = [-1,5,3,4,0]，输出 [-1,0,3,4,5]
// 示例 3：head = []，输出 []
//
// 思路：自顶向下归并排序（分治）。
// 快慢指针找中点并断成两半 → 递归排左右两半 → 回溯时（后序 = 合并点）
// 调 mergeTwoLists 归并，不新建节点。
// 时间 O(n log n)，递归栈 O(log n)。

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := slow.Next
	slow.Next = nil
	return mergeTwoLists(sortList(head), sortList(right))
}
