package problem_98

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1：递归1
func isValidBST(root *TreeNode) bool {
	_, _, isBst := proc(root)
	return isBst
}

func proc(root *TreeNode) (int, int, bool) { // 返回值为 当前子树最大和最小值、当前子树是否为BST
	if root == nil {
		return math.MinInt64, math.MaxInt64, true
	}
	lmax, lmin, lbst := proc(root.Left)
	rmax, rmin, rbst := proc(root.Right)
	max := maxInt(root.Val, maxInt(lmax, rmax))
	min := minInt(root.Val, minInt(lmin, rmin))
	bst := lbst && rbst && root.Val > lmax && root.Val < rmin
	return max, min, bst
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法2：二叉树中序遍历，检查是否为递增
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isFirst := true
	preVal := 0
	for root != nil {
		mostRight := root.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != root {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = root
				root = root.Left
				continue
			}
			if mostRight.Right == root {
				mostRight.Right = nil
			}
		}
		if !isFirst && root.Val <= preVal {
			return false
		}
		preVal = root.Val
		isFirst = false
		root = root.Right
	}
	return true
}

// 方法3：递归2
func isValidBST3(root *TreeNode) bool {
	return proc2(root, math.MinInt64, math.MaxInt64)
}

func proc2(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return proc2(root.Left, lower, root.Val) && proc2(root.Right, root.Val, upper)
}
