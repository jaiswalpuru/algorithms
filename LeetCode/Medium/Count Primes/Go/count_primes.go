package main

import "fmt"

func sieve(n int) []bool {
	is_prime := make([]bool, n)

	for i := 2; i < n; i++ {
		is_prime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if !is_prime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			is_prime[j] = false
		}
	}
	return is_prime
}

func count_prime(n int) int {
	cnt := 0
	is_prime := sieve(n)

	for i := 2; i < n; i++ {
		if is_prime[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(count_prime(10))
}
