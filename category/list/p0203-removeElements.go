package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeElements 203. 移除链表元素 https://leetcode.cn/problems/remove-linked-list-elements/description/
func removeElements(head *ListNode, val int) *ListNode {
	fakeNode := &ListNode{}
	fakeNode.Next = head
	currentNode := fakeNode
	for currentNode != nil && currentNode.Next != nil {
		if currentNode.Next.Val == val {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}
	return fakeNode.Next
}
