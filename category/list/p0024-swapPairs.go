package list

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// swapPairs 24. 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/description/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := head.Next
	first := head
	second := head.Next
	next := head
	var prev *ListNode
	for next != nil && next.Next != nil {
		next = second.Next
		second.Next = first
		first.Next = next
		if prev != nil {
			prev.Next = second
		}
		prev = first
		first = next
		if next != nil {
			second = next.Next
		}
	}
	return ans
}
