package problem0022

import "strings"

func generateParenthesis(n int) []string {
	res := []string{}
	p := ""
	var dfs func(string, int, int, int)
	dfs = func(p string, n int, l int, r int) {
		if l == n {
			res = append(res, p+strings.Repeat(")", l-r))
			return
		}

		if l < n {
			dfs(p+"(", n, l+1, r)
		}

		if r < l {
			dfs(p+")", n, l, r+1)
		}

		return
	}

	dfs(p, n, 0, 0)

	return res
}
