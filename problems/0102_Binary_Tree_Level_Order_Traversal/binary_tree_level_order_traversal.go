package problem0102

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	depth := depthOfTree(root, 0)
	res := make([][]int, depth)

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, k int) {
		if root.Left == nil && root.Right == nil {
			return
		}

		if root.Left != nil {
			res[k] = append(res[k], root.Left.Val)
			dfs(root.Left, k+1)
		}

		if root.Right != nil {
			res[k] = append(res[k], root.Right.Val)
			dfs(root.Right, k+1)
		}

		return
	}

	res[0] = []int{root.Val}
	dfs(root, 1)

	return res
}

func depthOfTree(root *TreeNode, k int) int {
	if root == nil {
		return k
	}

	left := depthOfTree(root.Left, k+1)
	right := depthOfTree(root.Right, k+1)

	if left < right {
		return right
	}

	return left
}
