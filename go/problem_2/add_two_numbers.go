package problem_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dh := &ListNode{}
	cur := dh
	more := 0
	for l1 != nil || l2 != nil || more > 0 {
		sum := 0
		if more > 0 {
			sum += more
		}
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		more = sum / 10
	}
	return dh.Next
}
