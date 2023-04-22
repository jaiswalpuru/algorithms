package main

import (
	"fmt"
	"sort"
)

func arithmetic_subarrays(arr []int, l, r []int) []bool {
	m := len(l)
	res := make([]bool, m)
	for i := 0; i < m; i++ {
		temp := []int{}
		for j := l[i]; j <= r[i]; j++ {
			temp = append(temp, arr[j])
		}
		if r[i]-l[i] < 1 {
			continue
		}
		sort.Ints(temp)
		res[i] = true
		for j := len(temp) - 1; j >= 2; j-- {
			if temp[j]-temp[j-1] != temp[j-1]-temp[j-2] {
				res[i] = false
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(arithmetic_subarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
}
