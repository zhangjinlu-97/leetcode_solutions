package problem_145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// morris遍历，在每个节点需要向右走时，逆序打印其左子树的右边界
// 注意：最后需要额外逆序打印整棵树的右边界
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0, 100)
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
			reverseAddRightEdge(cur.Left, &ret)
			cur = cur.Right
		}
		reverseAddRightEdge(root, &ret)
	}
	return ret
}

// 类似反转链表，先反转再逆序存入ret
func reverseAddRightEdge(root *TreeNode, ret *[]int) {
	cur := root
	var pre *TreeNode
	for cur != nil {
		next := cur.Right
		cur.Right = pre
		pre = cur
		cur = next
	}
	cur = pre
	pre = nil
	for cur != nil {
		*ret = append(*ret, cur.Val)
		next := cur.Right
		cur.Right = pre
		pre = cur
		cur = next
	}
}
