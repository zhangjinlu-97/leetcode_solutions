package problem_23

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dh := &ListNode{}
	cur := dh
	pq := make(PriorityQueue, 0, len(lists))
	for i := range lists {
		if lists[i] != nil { // 注意过滤掉nil
			pq = append(pq, lists[i])
		}
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		nd := heap.Pop(&pq).(*ListNode)
		cur.Next = nd
		cur = cur.Next
		if nd.Next != nil {
			heap.Push(&pq, nd.Next)
		}
	}
	return dh.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return x
}
