package problem_236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, _, ca := process(root, p, q)
	return ca
}

// 递归遍历二叉树，维护三个信息，子树中是否有p、子树中是否有q、当前的公共祖先节点
func process(root, p, q *TreeNode) (hasP, hasQ bool, ca *TreeNode) {
	if root == nil {
		return false, false, nil
	}
	lhp, lhq, lca := process(root.Left, p, q)
	rhp, rhq, rca := process(root.Right, p, q)
	hasP = lhp || rhp || root == p
	hasQ = lhq || rhq || root == q
	if lca != nil {
		ca = lca
		return
	}
	if rca != nil {
		ca = rca
		return
	}
	// 如果左右子树中都不存在公共祖先
	//则看p，q是否都存在与当前子树中
	// 如果是，则root节点即为公共祖先
	if hasP && hasQ {
		ca = root
	}
	return
}
