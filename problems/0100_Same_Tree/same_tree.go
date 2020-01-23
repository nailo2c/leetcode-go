package problem0100

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return false
}
