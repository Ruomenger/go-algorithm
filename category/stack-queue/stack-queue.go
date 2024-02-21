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

type MyStack struct {
	queue []int
}

func Constructor1() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (s *MyStack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	for ; n > 0; n-- {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}

}

func (s *MyStack) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

func (s *MyStack) Top() int {
	return s.queue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

// isValid 20. 有效的括号 https://leetcode.cn/problems/valid-parentheses/description/
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	stack := make([]int32, 0)
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')':
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}

	}
	return len(stack) == 0
}
