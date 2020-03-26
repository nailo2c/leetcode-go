package problem0094

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	res = helper(root, res)

	return res
}

func helper(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	if root.Left != nil {
		res = helper(root.Left, res)
	}

	res = append(res, root.Val)

	if root.Right != nil {
		res = helper(root.Right, res)
	}

	return res
}
