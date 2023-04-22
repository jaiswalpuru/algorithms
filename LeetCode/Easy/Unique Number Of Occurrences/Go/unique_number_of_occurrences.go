package main

import "fmt"

func unique_number_of_occurrences(arr []int) bool {
	mp := make(map[int]int)
	cnt := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		mp[arr[i]]++
	}
	for _, v := range mp {
		if cnt[v] > 0 {
			return false
		}
		cnt[v] = v
	}
	return true
}

func main() {
	fmt.Println(unique_number_of_occurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
