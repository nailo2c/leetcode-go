package problem0948

import "sort"

func bagOfTokensScore(tokens []int, P int) int {
	score := 0
	sort.Slice(tokens, func(i, j int) bool { return tokens[i] < tokens[j] })
	for len(tokens) > 0 && P >= tokens[0] {
		P = P - tokens[0]
		score++
		tokens = tokens[1:]

		if len(tokens) > 1 && P < tokens[0] {
			P = P + tokens[len(tokens)-1]
			score--
			tokens = tokens[:len(tokens)-1]
		}
	}

	return score
}
