package main

import "fmt"

func max_number_of_k_sum_pairs(arr []int, k int) int {
	hash_map := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		hash_map[arr[i]]++
	}
	cnt := 0
	for i := 0; i < n; i++ {
		val := k - arr[i]
		if hash_map[arr[i]] > 0 && hash_map[val] > 0 {
			if val == arr[i] && hash_map[val] < 2 {
				continue
			}
			hash_map[arr[i]]--
			hash_map[val]--
			cnt += 1
		}
	}
	return cnt
}

func main() {
	fmt.Println(max_number_of_k_sum_pairs([]int{1, 2, 3, 4}, 5))
}
