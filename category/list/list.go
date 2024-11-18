package list

type ListNode struct {
	Val  int
	Next *ListNode
}

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

// 707. 设计链表 https://leetcode.cn/problems/design-linked-list/description/

type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	val  int
	next *Node
	pre  *Node
}

func Constructor() MyLinkedList {
	dummy := &Node{
		val:  -1,
		next: nil,
		pre:  nil,
	}
	dummy.next = dummy
	dummy.pre = dummy
	return MyLinkedList{
		dummy: dummy,
	}
}

func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.next
	for head != this.dummy && index > 0 {
		index--
		head = head.next
	}
	if index != 0 {
		return -1
	}
	return head.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		val:  val,
		next: dummy.next,
		pre:  dummy,
	}

	dummy.next.pre = node
	dummy.next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	node := &Node{
		val:  val,
		next: dummy,
		pre:  dummy.pre,
	}
	dummy.pre.next = node
	dummy.pre = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.dummy.next
	for head != this.dummy && index > 0 {
		head = head.next
		index--
	}
	if index > 0 {
		return
	}
	node := &Node{
		val:  val,
		next: head,
		pre:  head.pre,
	}
	head.pre.next = node
	head.pre = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.dummy.next == this.dummy {
		return
	}
	head := this.dummy.next
	for head.next != this.dummy && index > 0 {
		head = head.next
		index--
	}
	if index == 0 {
		head.next.pre = head.pre
		head.pre.next = head.next
	}
}

// mergeTwoLists 21. 合并两个有序链表
// level: 简单
// tag: 链表
// https://leetcode.cn/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = cur.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}
