package problem0437

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	k := 0
	var dfs func(*TreeNode, int, int)
	dfs = func(root *TreeNode, sum int, accumulate int) {
		if root == nil {
			return
		}

		accumulate += root.Val
		if accumulate == sum {
			k++
		}
		dfs(root.Left, sum, accumulate)
		dfs(root.Right, sum, accumulate)

		return
	}

	var search func(*TreeNode, *TreeNode, int)
	search = func(t1 *TreeNode, t2 *TreeNode, sum int) {
		dfs(t1, sum, 0)
		dfs(t2, sum, 0)

		if t1 != nil {
			search(t1.Left, t1.Right, sum)
		}

		if t2 != nil {
			search(t2.Left, t2.Right, sum)
		}

		return
	}

	dfs(root, sum, 0)
	search(root.Left, root.Right, sum)

	return k
}
