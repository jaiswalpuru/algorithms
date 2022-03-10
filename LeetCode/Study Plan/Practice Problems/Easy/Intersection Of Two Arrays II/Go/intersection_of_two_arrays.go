package main

import "fmt"

func intersection_of_two_arrays(a1, a2 []int) []int {
	n, m := len(a1), len(a2)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[a1[i]]++
	}

	res := []int{}
	for i := 0; i < m; i++ {
		if v, ok := mp[a2[i]]; ok && v > 0 {
			res = append(res, a2[i])
			mp[a2[i]]--
		}
	}
	return res
}

func main() {
	fmt.Println(intersection_of_two_arrays([]int{1, 2, 2, 1}, []int{2, 2}))
}
