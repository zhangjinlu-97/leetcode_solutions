package problem_25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dh := &ListNode{Next: head}
	pre, cur := dh, head
	for cur != nil {
		end := cur
		for i := 0; i < k-1 && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		next := end.Next
		end.Next = nil
		reverseList(cur)
		pre.Next = end
		pre = cur
		cur.Next = next
		cur = next
	}
	return dh.Next
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
