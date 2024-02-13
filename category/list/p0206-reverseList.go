package list

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// reverseList 206. 反转链表 https://leetcode.cn/problems/reverse-linked-list/description/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head.Next
	head.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = head
		head.Next = cur.Next.Next
		head = cur
		cur = next
	}
	return head
}
