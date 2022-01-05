package problem_129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//DFS，维护当前路径上的数，当到达叶子节点时，在结果中加上当前路径上的数
func sumNumbers(root *TreeNode) int {
	sum := 0
	process(root, 0, &sum)
	return sum
}

func process(root *TreeNode, num int, sum *int) {
	if root == nil {
		return
	}
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += num
	}
	process(root.Left, num, sum)
	process(root.Right, num, sum)
}
