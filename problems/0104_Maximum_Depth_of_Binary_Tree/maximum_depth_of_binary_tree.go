package problem0104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var k int
	return depth(root, k)
}

func depth(root *TreeNode, k int) int {
	if root != nil {
		k++
		return max(depth(root.Left, k), depth(root.Right, k))
	}

	return k
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
