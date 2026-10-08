package lc

// 141. 环形链表（简单）
//
// 给你一个链表的表头节点 head，判断该链表中是否存在环。
// 存在环：如果链表中某个节点指向了该节点之前的节点，再次指向该节点则链表形成环。
// 提示：链表节点只会通过 Next 指针指向下一个节点。
// 进阶：能否用 O(1) 空间复杂度解决？
//
// 示例 1：head = [3,2,0,-4]，输出 true
//        链表中存在环，故返回 true。
// 示例 2：head = [1,2]，输出 false
//        链表中不存在环，故返回 false。
//
// 提示：思路为快慢指针（速度差），fast 走 2 步、slow 走 1 步。

func hasCycle(head *ListNode) bool {
   if head==nil{
	return false
   }
	fast,slow:=head,head
   for fast.Next!=nil&&fast.Next.Next!=nil{
    fast=fast.Next.Next
    slow=slow.Next
		if fast==slow{
			return true
		}
	}
   }
   return false
}
