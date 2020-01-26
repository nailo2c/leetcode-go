package problem0204

func countPrimes(n int) int {
	primes := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if !primes[i] {
			for j := i; i*j < n; j++ {
				primes[i*j] = true
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if !primes[i] {
			count++
		}
	}
	return count
}
