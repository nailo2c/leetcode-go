package problem0139

func wordBreak(s string, wordDict []string) bool {
	mem := make([]bool, len(s)+1)
	mem[0] = true

	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if mem[j] && dict[s[j:i]] {
				mem[i] = true
				break
			}
		}
	}

	return mem[len(s)]
}
