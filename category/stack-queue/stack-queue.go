package stack_queue

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

type MyQueue struct {
	stackIn  *Stack
	stackOut *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  &Stack{},
		stackOut: &Stack{},
	}
}

func (q *MyQueue) Push(x int) {
	q.stackIn.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.stackOut.Empty() {
		for !q.stackIn.Empty() {
			val := q.stackIn.Pop()
			q.stackOut.Push(val)
		}
	}
	return q.stackOut.Pop()
}

func (q *MyQueue) Peek() int {
	if q.stackOut.Empty() {
		for !q.stackIn.Empty() {
			val := q.stackIn.Pop()
			q.stackOut.Push(val)
		}
	}
	return q.stackOut.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.stackIn.Empty() && q.stackOut.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
