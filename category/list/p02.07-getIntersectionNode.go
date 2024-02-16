package list

// getIntersectionNode 面试题 02.07. 链表相交 https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	sub := 0
	var fast, slow *ListNode
	if lenA > lenB {
		sub = lenA - lenB
		fast, slow = headA, headB
	} else {
		sub = lenB - lenA
		fast, slow = headB, headA
	}
	for sub > 0 {
		fast = fast.Next
		sub--
	}
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
