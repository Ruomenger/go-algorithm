package list

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// removeNthFromEnd 19. 删除链表的倒数第 N 个结点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	fast := head
	for ; n > 0; n-- {
		fast = fast.Next
	}
	slow := dummy
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}
