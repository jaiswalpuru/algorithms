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

	next_set := []int{}
	last_num := num % 10
	next_set = append(next_set, last_num+k)
	if k != 0 {
		next_set = append(next_set, last_num-k)
	}
	for i := 0; i < len(next_set); i++ {
		if next_set[i] >= 0 && next_set[i] < 10 {
			temp := num*10 + next_set[i]
			dfs(n-1, k, res, temp)
		}
	}
}

//--------------------using dfs-------------------------

//--------------------using bfs-------------------------
func number_with_same_consecutive_diff_bfs(n, k int) []int {
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 1; i < n; i++ { //level
		next_q := []int{}
		for _, v := range q {
			last_num := v % 10
			next_set := []int{}
			next_set = append(next_set, last_num+k)
			if k != 0 {
				next_set = append(next_set, last_num-k)
			}
			for _, v2 := range next_set {
				if v2 >= 0 && v2 < 10 {
					new_num := v*10 + v2
					next_q = append(next_q, new_num)
				}
			}
		}
		q = next_q
	}
	return q
}

//--------------------using dfs-------------------------

func main() {
	fmt.Println(number_with_same_consecutive_diff_bfs(3, 7))
}
