package lc

// 142. 环形链表 II（中等）
//
// 给定一个链表的表头节点 head，返回链表开始入环的第一个节点。
// 如果链表没有环，则返回 nil。
// 进阶：能否用 O(1) 空间复杂度解决？
//
// 示例 1：head = [3,2,0,-4]，输出 2
//        链表中存在环，故返回入环节点 2。
// 示例 2：head = [1,2]，输出 nil
//        链表中不存在环，故返回 nil。
//

// 设 head 到入环点的距离为 a，环长为 c，相遇点到入环点的距离为 b
// slow 走 t 步，fast 走 2t 步
// 相遇时：
//	a + b = t
//	2t = a + b + k*c (k >= 1，因为 fast 多走的那段必须绕整数圈)
// 两式关联得：a = k*c - b   (k >= 1)
//
// slow 从相遇点出发，每 c-b 步经过一次入环点，之后每多绕一圈(c步)再经过一次
//	第 n 次经过入环点的时刻 = n*c - b   (n >= 1)
// 因为 a = k*c - b，所以 head 到入环点的距离 a 恰好是 slow 第 k 次经过入环点的时刻
// 于是 p 从 head 走 a 步到达入环点时，slow 正好第 k 次经过入环点，两者必相遇，相遇就在入环点


func detectCycle(head *ListNode) *ListNode {
	if head==nil{
		return false
	}
  fast,slow:=head,head
 for fast.Next!=nil&&fast.Next!=nil{
	fast=fast.Next.Next
	slow=slow.Next
	if fast==slow{
		p:=head
		for p!=slow {
			p=p.Next
			slow=slow.Next
		}
		return p
	}
 }
    return nil
}
