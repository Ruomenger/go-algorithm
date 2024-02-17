package go_algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	first := true
	fast, slow := head, head
	for fast != slow || first {
		first = false
		slow = slow.Next
		fast = fast.Next
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next
	}
	return true
}
