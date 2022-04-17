package main

import "fmt"

func minimum_rounds_to_complete_all_tasks(arr []int) int {
	hash_map := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		hash_map[arr[i]]++
	}

	res := 0
	for _, v := range hash_map {
		if v == 1 {
			return -1
		}
		res += (v + 2) / 3
	}
	return res
}

func main() {
	fmt.Println(minimum_rounds_to_complete_all_tasks([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
}
