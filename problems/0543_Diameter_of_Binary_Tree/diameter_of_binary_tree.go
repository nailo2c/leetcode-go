package problem0543

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, k int) int {
		if root == nil {
			return k
		}
		return max(dfs(root.Left, k+1), dfs(root.Right, k+1))
	}

	rootSum := dfs(root.Left, 0) + dfs(root.Right, 0)

	return maxTrip(rootSum, diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxTrip(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}

	if b >= a && b >= c {
		return b
	}

	return c
}
