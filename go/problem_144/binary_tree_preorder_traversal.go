package problem_144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris遍历 时间: O(N); 空间: O(1)
// 当一个节点没有左树时，也要加入ret，否则会漏掉
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0, 100)
	if root != nil {
		for root != nil {
			mostRight := root.Left
			if mostRight != nil {
				for mostRight.Right != nil && mostRight.Right != root {
					mostRight = mostRight.Right
				}
				if mostRight.Right == nil {
					ret = append(ret, root.Val)
					mostRight.Right = root
					root = root.Left
					continue
				}
				if mostRight.Right == root {
					mostRight.Right = nil
				}
			}
			root = root.Right
		}
	}
	return ret
}
