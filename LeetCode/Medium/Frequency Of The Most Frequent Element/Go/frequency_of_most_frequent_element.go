package main

import (
	"fmt"
	"sort"
)

//---------------------------brute force---------------------------
func frequency_of_most_frequent_element(arr []int, k int) int {
	sort.Ints(arr)
	n := len(arr)
	cnt := 1
	for i := n - 1; i >= 0; i-- {
		curr := 1
		val := arr[i]
		t := k
		for j := i - 1; j >= 0; j-- {
			if (val - arr[j]) >= 0 {
				t -= (val - arr[j])
				if t < 0 {
					break
				}
				curr++
			}
		}
		cnt = max(cnt, curr)
	}
	return cnt
}

//---------------------------brute force---------------------------

//---------------------------efficient soln---------------------------
func frequency_of_most_frequent_element_eff(arr []int, k int) int {
	n := len(arr)
	sort.Ints(arr)
	res := 1
	l, sum := 0, 0
	for r := 0; r < n; r++ {
		sum += arr[r]
		for arr[r]*(r-l+1)-sum > k {
			sum -= arr[l]
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

//---------------------------efficient soln---------------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(frequency_of_most_frequent_element_eff([]int{1, 2, 4}, 5))
}
