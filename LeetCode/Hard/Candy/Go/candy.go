package main

import (
	"fmt"
	"math"
)

//----------------brute force------------------------
func candy(arr []int) int {
	mp := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		mp[i] = 1
	}
	has_changed := true
	for has_changed {
		has_changed = false
		for i := 0; i < n; i++ {
			if i != n-1 && arr[i] > arr[i+1] && mp[i] <= mp[i+1] {
				mp[i] = mp[i+1] + 1
				has_changed = true
			}
			if i > 0 && arr[i] > arr[i-1] && mp[i] <= mp[i-1] {
				mp[i] = mp[i-1] + 1
				has_changed = true
			}
		}
	}
	cnt := 0
	for _, v := range mp {
		cnt += v
	}
	return cnt
}

//----------------brute force------------------------

//----------------efficient approach------------------------
func candy_eff(arr []int) int {
	n := len(arr)
	l_r := make([]int, n)
	r_l := make([]int, n)
	for i := 0; i < n; i++ {
		l_r[i] = 1
		r_l[i] = 1
	}

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			l_r[i] = l_r[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			r_l[i] = r_l[i+1] + 1
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		cnt += int(math.Max(float64(l_r[i]), float64(r_l[i])))
	}
	return cnt
}

//----------------efficient approach------------------------

func main() {
	fmt.Println(candy_eff([]int{1, 0, 2}))
}
