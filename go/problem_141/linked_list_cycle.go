package problem_141

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func hasCycle(head *ListNode) bool {
	fa, sl := head, head
	for fa != nil && fa.Next != nil {
		fa = fa.Next.Next
		sl = sl.Next
		if fa == sl {
			return true
		}
	}
	return false
}
