package problem0114

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil && (root.Left.Left != nil || root.Left.Right != nil) {
		flatten(root.Left)
	}

	if root.Right != nil && (root.Right.Left != nil || root.Right.Right != nil) {
		flatten(root.Right)
	}

	if root.Left != nil && root.Right != nil {
		tmp := root.Right
		root.Right = root.Left
		root := rootRightTail(root.Right)
		root.Right = tmp
	}

	if root.Left != nil && root.Right == nil {
		root.Right = root.Left
	}

	root.Left = nil
	return
}

func rootRightTail(root *TreeNode) *TreeNode {
	if root.Right == nil {
		return root
	}
	return rootRightTail(root.Right)
}
