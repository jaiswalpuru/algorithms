package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//--------------------illegal method but works :) ------------------
func longest_consecutive_sequence(arr []int) int {
	sort.Ints(arr)
	res := 0
	t := 1
	n := len(arr)
	if n == 0 {
		return 0
	}
	for i := 0; i < n-1; i++ {
		if arr[i+1] == arr[i] {
			continue
		}
		if arr[i+1]-arr[i] == 1 {
			t++
		} else {
			res = max(res, t)
			t = 1
		}
	}
	res = max(res, t)
	return res
}

//--------------------illegal method but works :) ------------------

//-------------------legal way to solve this problem as asked to do in O(n) the above method takes O(nlogn)------------
func longest_consecutive_sequence_eff(arr []int) int {
	mp := make(map[int]struct{})
	n := len(arr)
	for i := 0; i < n; i++ {
		mp[arr[i]] = struct{}{}
	}

	res := 0
	for i := 0; i < n; i++ {
		if _, ok := mp[arr[i]-1]; !ok {
			curr := arr[i]
			cnt := 0
			_, ok := mp[curr]
			for ok {
				cnt++
				curr += 1
				_, ok = mp[curr]
			}
			res = max(res, cnt)
		}
	}
	return res
}

//-------------------legal way to solve this problem as asked to do in O(n) the above method takes O(nlogn)------------

func main() {
	fmt.Println(longest_consecutive_sequence_eff([]int{100, 4, 200, 1, 3, 2}))
}
