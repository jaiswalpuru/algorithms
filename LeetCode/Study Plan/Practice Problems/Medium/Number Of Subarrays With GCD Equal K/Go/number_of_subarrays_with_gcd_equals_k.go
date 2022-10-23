package main

import "fmt"

func number_of_subarrays_with_gcd_equals_k(arr []int, k int) int {
	n := len(arr)
	cnt := 0
	for i := 0; i < n; i++ {
		curr := 0
		for j := i; j < n; j++ {
			curr = gcd(curr, arr[j])
			if curr == k {
				cnt++
			}
		}
	}
	return cnt
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Println(number_of_subarrays_with_gcd_equals_k([]int{9, 3, 1, 2, 6, 3}, 3))
}
