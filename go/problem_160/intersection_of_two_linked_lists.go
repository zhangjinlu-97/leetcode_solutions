package problem_160

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 分别遍历两个链表,获得链表长度
// 2. 长的链表先走 两链表长度之差步
// 3. 两链表同时向前，在遇到nil之前，如果发现相同的节点，表示两链表相交
// 注意：如果两链表相交，一定共用同一段链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := getListLen(headA), getListLen(headB)
	var long, short *ListNode
	var n int
	if lenA > lenB {
		long = headA
		short = headB
		n = lenA - lenB
	} else {
		long = headB
		short = headA
		n = lenB - lenA
	}
	for i := 0; i < n; i++ {
		long = long.Next
	}
	for long != nil && long != short {
		long = long.Next
		short = short.Next
	}
	return long
}

func getListLen(head *ListNode) int {
	n := 0
	for head != nil {
		head = head.Next
		n++
	}
	return n
}
