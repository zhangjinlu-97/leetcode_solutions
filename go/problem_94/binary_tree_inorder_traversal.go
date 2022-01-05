package problem_94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris遍历第二次到达某节点时，抓答案
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root != nil {
		cur := root
		for cur != nil {
			mostRight := cur.Left
			if mostRight != nil {
				for mostRight.Right != nil && mostRight.Right != cur {
					mostRight = mostRight.Right
				}
				if mostRight.Right == nil {
					mostRight.Right = cur
					cur = cur.Left
					continue
				}
				if mostRight.Right == cur {
					mostRight.Right = nil
				}
			}
			ret = append(ret, cur.Val)
			cur = cur.Right
		}
	}
	return ret
}
