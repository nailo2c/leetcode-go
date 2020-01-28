package problem0257

import "strconv"

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	res := []string{}

	var dfs func(string, *TreeNode)
	dfs = func(pre string, root *TreeNode) {
		if pre == "" {
			pre = strconv.Itoa(root.Val)
		} else {
			pre += "->" + strconv.Itoa(root.Val)
		}

		if root.Left != nil {
			dfs(pre, root.Left)
		}

		if root.Right != nil {
			dfs(pre, root.Right)
		}

		if root.Left == nil && root.Right == nil {
			res = append(res, pre)
		}

	}

	dfs("", root)

	return res
}
