package problem_143

import (
	"fmt"
	"testing"
)

func TestReorderList(t *testing.T) {
	n5 := &ListNode{Val: 5, Next: nil}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	printList(n1)
	reorderList(n1)
	printList(n1)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
