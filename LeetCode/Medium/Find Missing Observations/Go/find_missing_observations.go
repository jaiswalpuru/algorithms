package main

import "fmt"

func find_missing_observations(roles []int, mean, n int) []int {
	role_sum := 0
	m := len(roles)
	for i := 0; i < m; i++ {
		role_sum += roles[i]
	}

	missing := mean*(m+n) - role_sum

	if missing > n*6 || missing < n {
		return []int{}
	}

	res := []int{}
	for missing > 0 {
		val := missing / n
		res = append(res, val)
		missing -= val
		n--
	}
	return res
}

func main() {
	fmt.Println(find_missing_observations([]int{3, 2, 4, 3}, 4, 2))
}
