package main

import (
	"fmt"
)

//--------------------------brute force------------------------------------
func minimum_number_of_days_to_make_m_bouquets(arr []int, m, k int) int {
	n := len(arr)
	max_flower := 0
	for i := 0; i < n; i++ {
		max_flower = max(max_flower, arr[i])
	}
	for i := 1; i <= max_flower; i++ {
		cnt := 0
		t := m
		for j := 0; j < n; j++ {
			if arr[j] <= i {
				cnt++
				if cnt == k {
					t--
					cnt = 0
				}
			} else {
				cnt = 0
			}
			if t == 0 {
				return i
			}
		}
	}
	return -1
}

//--------------------------brute force------------------------------------

//--------------------------efficient approach------------------------------------
func can_make_enough(arr []int, k int, day int) int {
	cnt := 0
	b := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > day {
			cnt = 0
		} else {
			cnt++
			if cnt >= k {
				b++
				cnt = 0
			}
		}
	}
	return b
}

func minimum_number_of_days_to_make_m_bouquets_eff(arr []int, m, k int) int {
	if m*k > len(arr) {
		return -1
	}
	min_val, max_val := arr[0], arr[0]
	for i := 0; i < len(arr); i++ {
		min_val = min(min_val, arr[i])
		max_val = max(max_val, arr[i])
	}
	l, r := min_val, max_val
	for l < r {
		mid := (l + r) / 2
		v := can_make_enough(arr, k, mid)
		if v < m {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

//--------------------------efficient approach------------------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimum_number_of_days_to_make_m_bouquets_eff([]int{1, 10, 3, 10, 2}, 3, 1))
}
