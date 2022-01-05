package problem_199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root != nil {
		q := &queue{}
		q.offer(root)
		curEnd := root
		nextEnd := (*TreeNode)(nil)
		for !q.isEmpty() {
			cur := q.poll()
			if cur.Left != nil {
				q.offer(cur.Left)
				nextEnd = cur.Left
			}
			if cur.Right != nil {
				q.offer(cur.Right)
				nextEnd = cur.Right
			}
			if cur == curEnd {
				ans = append(ans, cur.Val)
				curEnd = nextEnd
			}
		}
	}
	return ans
}

type queue []*TreeNode

func (q *queue) offer(tn *TreeNode) {
	*q = append(*q, tn)
}

func (q *queue) poll() *TreeNode {
	old := *q
	tn := old[0]
	*q = old[1:]
	old[0] = nil
	return tn
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}
