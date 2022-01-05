package problem_103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//宽度优先遍历
//注意，root为nil时直接返回；队列别写成了栈
//end记录当前层的结尾
//nextEnd记录下一层结尾
//遍历时只要有子节点，就更新nextEnd
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0, 16)
	queue := func(q Queue) *Queue { return &q }(make([]*TreeNode, 0, 16))
	var end, nextEnd *TreeNode
	fromLeft := true
	end = root
	queue.offer(root)
	levelList := make([]int, 0)
	for !queue.isEmpty() {
		node := queue.poll()
		if fromLeft {
			levelList = append(levelList, node.Val)
		} else {
			levelList = append([]int{node.Val}, levelList...)
		}
		if node.Left != nil {
			nextEnd = node.Left
			queue.offer(node.Left)
		}
		if node.Right != nil {
			nextEnd = node.Right
			queue.offer(node.Right)
		}
		if end == node {
			res = append(res, levelList)
			end = nextEnd
			levelList = make([]int, 0)
			fromLeft = !fromLeft
		}
	}
	return res
}

type Queue []*TreeNode

func (q *Queue) offer(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) poll() *TreeNode {
	old := *q
	x := old[0]
	*q = old[1:]
	return x
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}
