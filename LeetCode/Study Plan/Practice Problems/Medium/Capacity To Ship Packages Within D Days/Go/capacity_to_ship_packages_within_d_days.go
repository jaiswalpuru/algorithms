package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func is_possible(wt []int, d, capa int) bool {
	curr_wt := 0
	total_day := 1
	for i := 0; i < len(wt); i++ {
		curr_wt += wt[i]
		if curr_wt > capa {
			curr_wt = wt[i]
			total_day += 1
		}
	}
	return total_day <= d
}

func capacity_to_ship_packages_within_d_days(w []int, d int) int {
	total_wt := 0
	max_wt := 0
	for i := 0; i < len(w); i++ {
		total_wt += w[i]
		max_wt = max(max_wt, w[i])
	}

	l, h := max_wt, total_wt
	for l < h {
		mid := (l + h) / 2
		if is_possible(w, d, mid) {
			h = mid
		} else {
			l = mid + 1
		}
	}
	return h
}

func main() {
	fmt.Println(capacity_to_ship_packages_within_d_days([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
