package problem_155

// MinStack 两个栈
type MinStack struct {
	head    *node
	minNode *node
}

type node struct {
	val  int
	next *node
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	if ms.minNode == nil {
		ms.minNode = &node{val: val}
	} else {
		n := &node{}
		if val < ms.minNode.val {
			n.val = val
		} else {
			n.val = ms.minNode.val
		}
		n.next = ms.minNode
		ms.minNode = n
	}
	n := &node{val: val}
	if ms.head == nil {
		ms.head = n
	} else {
		n.next = ms.head
		ms.head = n
	}
}

func (ms *MinStack) Pop() {
	ms.head = ms.head.next
	ms.minNode = ms.minNode.next
}

func (ms *MinStack) Top() int {
	return ms.head.val
}

func (ms *MinStack) GetMin() int {
	return ms.minNode.val
}
