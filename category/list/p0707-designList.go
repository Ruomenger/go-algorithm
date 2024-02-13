package list

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

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
