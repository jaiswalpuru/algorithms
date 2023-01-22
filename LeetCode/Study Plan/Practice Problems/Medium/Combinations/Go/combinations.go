package main

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}
	backtrack(n, k, 1, []int{}, &res)
	return res
}

func backtrack(n, k, start int, temp []int, res *[][]int) {
	if k == 0 {
		*res = append(*res, append([]int{}, temp...))
		return
	}
	for i := start; i <= n; i++ {
		temp = append(temp, i)
		backtrack(n, k-1, i+1, temp, res)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(combination(4, 2))
}
