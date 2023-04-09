package prmie

func Eratosthenes(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	res := make([]int, 0, n)
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			res = append(res, i)
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	return res
}

func Euler(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	res := make([]int, 0, n)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			res = append(res, i)
		}
		for j := 0; j < len(res) && i*res[j] <= n; j++ {
			isPrime[i*res[j]] = false
			if i%res[j] == 0 {
				break
			}
		}
	}

	return res
}
