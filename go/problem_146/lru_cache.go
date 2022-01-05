package problem_146

// hash表 + 双向链表
type LRUCache struct {
	nodeMap    map[int]*node
	linkedList *LinkedList
	size       int
	cap        int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		nodeMap:    map[int]*node{},
		linkedList: NewLinkedList(),
		size:       0,
		cap:        capacity,
	}
}

func (lc *LRUCache) Get(key int) int {
	if n, ok := lc.nodeMap[key]; ok {
		lc.linkedList.MoveToLast(n)
		return n.value
	}
	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	if n, ok := lc.nodeMap[key]; ok {
		n.value = value
		lc.linkedList.MoveToLast(n)
		return
	}
	if lc.size == lc.cap {
		rn := lc.linkedList.RemoveFirst()
		delete(lc.nodeMap, rn.key)
		lc.size--
	}
	newNode := &node{key: key, value: value}
	lc.nodeMap[key] = newNode
	lc.linkedList.Add(newNode)
	lc.size++
}

// 实现双向链表
type LinkedList struct {
	head *node
	tail *node
}

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

func NewLinkedList() *LinkedList {
	l := &LinkedList{
		head: new(node),
		tail: new(node),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LinkedList) Add(n *node) {
	n.prev = l.tail.prev
	n.next = l.tail
	n.prev.next = n
	n.next.prev = n
}

func (l *LinkedList) MoveToLast(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	l.Add(n)
}

func (l *LinkedList) RemoveFirst() *node {
	n := l.head.next
	n.next.prev = l.head
	l.head.next = n.next
	return n
}
