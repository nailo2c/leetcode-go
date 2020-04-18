package problem0105

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pStart int

func buildTree(preorder []int, inorder []int) *TreeNode {
	pStart = 0
	return helper(preorder, inorder, 0, len(inorder)-1)
}

func helper(preorder, inorder []int, iStart, iEnd int) *TreeNode {
	if iStart > iEnd {
		return nil
	}

	node := &TreeNode{Val: preorder[pStart]}
	pStart++

	var idx int
	for idx = iStart; idx <= iEnd; idx++ {
		if inorder[idx] == node.Val {
			break
		}
	}

	node.Left = helper(preorder, inorder, iStart, idx-1)
	node.Right = helper(preorder, inorder, idx+1, iEnd)

	return node
}
