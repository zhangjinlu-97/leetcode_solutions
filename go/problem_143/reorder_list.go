package problem_143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	sl, fa := head, head
	for fa != nil && fa.Next != nil {
		sl = sl.Next
		fa = fa.Next.Next
	}
	tail := reverseList(sl) // 反转后半段链表，并拿到尾部节点
	cur := (*ListNode)(nil)
	for head != nil && tail != nil && head != tail {
		if cur != nil {
			cur.Next = head
		}
		next := head.Next
		head.Next = tail
		head = next
		cur = tail
		tail = tail.Next
	}
	if head == tail && cur != nil { // 节点个数为奇数时，最终会来到同一节点
		cur.Next = head
	}
}

func reverseList(head *ListNode) *ListNode {
	pre := (*ListNode)(nil)
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
