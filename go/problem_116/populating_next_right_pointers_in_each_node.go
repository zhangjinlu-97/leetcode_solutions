package problem_116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Offer(node *Node) {
	q.size++
	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.Next = node
	q.tail = node
}

func (q *Queue) Poll() *Node {
	q.size--
	x := q.head
	q.head = x.Next
	x.Next = nil
	return x
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

// 使用题目中的node，实现一个单链表队列；从而使额外空间为O(1)
// 层序遍历，每层的节点依次连接即可
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	que := &Queue{}
	que.Offer(root)
	for !que.IsEmpty() {
		pre := (*Node)(nil)
		size := que.Size()
		for i := 0; i < size; i++ {
			cur := que.Poll()
			if cur.Left != nil {
				que.Offer(cur.Left)
			}
			if cur.Right != nil {
				que.Offer(cur.Right)
			}
			if pre != nil {
				pre.Next = cur
			}
			pre = cur
		}
	}
	return root
}
