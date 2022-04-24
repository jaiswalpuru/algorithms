package main

import "fmt"
import "sort"

func intersection(arr [][]int) []int {
	n := len(arr)
	res := []int{}
	hash_map := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < len(arr[i]); j++ {
			hash_map[arr[i][j]]++
		}
	}

	for k, v := range hash_map {
		if v == n {
			res = append(res, k)
		}
	}

	sort.Ints(res)
	return res
}

func main() {
	fmt.Println(intersection([][]int{
		{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6},
	}))
}
