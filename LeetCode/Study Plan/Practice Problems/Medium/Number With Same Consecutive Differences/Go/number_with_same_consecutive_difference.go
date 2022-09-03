package main

import "fmt"

//--------------------using dfs-------------------------
func number_with_same_consecutive_diff_dfs(n, k int) []int {
	res := []int{}
	for num := 1; num < 10; num++ {
		dfs(n-1, k, &res, num)
	}
	return res
}

func dfs(n int, k int, res *[]int, num int) {
	if n == 0 {
		*res = append(*res, num)
		return
	}

	last_num := num % 10
	if last_num+k < 10 {
		dfs(n-1, k, res, num*10+last_num+k)
	}
	if k != 0 && last_num-k >= 0 {
		dfs(n-1, k, res, num*10+last_num-k)
	}
}

//--------------------using dfs-------------------------

//--------------------using bfs-------------------------
func number_with_same_consecutive_diff_bfs(n, k int) []int {
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 1; i < n; i++ {
		next_q := []int{}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]
			last_num := curr % 10
			if last_num+k < 10 {
				next_q = append(next_q, curr*10+last_num+k)
			}
			if k != 0 && last_num-k >= 0 {
				next_q = append(next_q, curr*10+last_num-k)
			}
		}
		q = next_q
	}
	return q
}

//--------------------using bfs-------------------------

func main() {
	fmt.Println(number_with_same_consecutive_diff_bfs(3, 7))
}
