package problem_101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 两二叉树对称条件：
// 1.它们的两个根结点具有相同的值
// 2.每个树的右子树都与另一个树的左子树镜像对称
func isSymmetric(root *TreeNode) bool {
	return proc(root, root)
}

func proc(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && proc(left.Right, right.Left) && proc(left.Left, right.Right)
}
