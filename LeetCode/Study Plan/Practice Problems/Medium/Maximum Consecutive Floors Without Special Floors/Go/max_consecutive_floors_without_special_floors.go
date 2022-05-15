package main

import "sort"
import "fmt"

func max_consecutive_floors_after_removing_special_floors(bottom int, top int, special []int) int {
	sort.Ints(special)
	res := 0
	res = max(res, special[0]-bottom)
	for i := 1; i < len(special); i++ {
		res = max(res, special[i]-special[i-1]-1)
	}
	res = max(res, top-special[len(special)-1])
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max_consecutive_floors_after_removing_special_floors(2, 9, []int{4, 6}))
}
