package main

import "fmt"

func intersection_of_two_arrays(a, b []int) []int {
	ma := make(map[int]int)
	for i := 0; i < len(a); i++ {
		ma[a[i]]++
	}
	res := make(map[int]bool)
	for i := 0; i < len(b); i++ {
		if ma[b[i]] != 0 {
			res[b[i]] = true
			ma[b[i]]--
		}
	}
	ans := []int{}
	for k := range res {
		ans = append(ans, k)
	}
	return ans
}

func main() {
	fmt.Println(intersection_of_two_arrays([]int{1, 2, 2, 1}, []int{2, 2}))
}
