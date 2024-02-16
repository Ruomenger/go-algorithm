package list

// detectCycle 142. 环形链表 II https://leetcode.cn/problems/linked-list-cycle-ii
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	first := true
	fast, slow := head, head
	for fast != slow || first {
		first = false
		slow = slow.Next
		fast = fast.Next
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// hasCycle 141. 环形链表 https://leetcode.cn/problems/linked-list-cycle
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
