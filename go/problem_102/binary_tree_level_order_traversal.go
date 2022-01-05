package problem_102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	queue := &Queue{}
	queue.offer(root)
	curEnd, nextEnd := root, &TreeNode{}
	var levelRet []int
	for !queue.isEmpty() {
		cur := queue.poll()
		levelRet = append(levelRet, cur.Val)
		if cur.Left != nil {
			queue.offer(cur.Left)
			nextEnd = cur.Left
		}
		if cur.Right != nil {
			queue.offer(cur.Right)
			nextEnd = cur.Right
		}
		if cur == curEnd {
			curEnd = nextEnd
			ret = append(ret, levelRet)
			levelRet = []int{}
		}
	}
	return ret
}

type Queue []*TreeNode

func (queue *Queue) offer(node *TreeNode) {
	*queue = append(*queue, node)
}

func (queue *Queue) poll() *TreeNode {
	x := (*queue)[0]
	*queue = (*queue)[1:]
	return x
}

func (queue *Queue) isEmpty() bool {
	return len(*queue) == 0
}
