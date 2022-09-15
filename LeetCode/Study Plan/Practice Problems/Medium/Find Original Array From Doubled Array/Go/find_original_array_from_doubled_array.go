package main

import (
	"fmt"
	"sort"
)

func find(changed []int) []int {
	res := []int{}
	n := len(changed)
	if n%2 != 0 {
		return res
	}
	hash_map := make(map[int]int)
	for i := 0; i < n; i++ {
		hash_map[changed[i]]++
	}
	sort.Ints(changed)
	for i := 0; i < n; i++ {
		val := changed[i]
		if _, ok := hash_map[val]; ok {
			hash_map[val]--
			if hash_map[val] == 0 {
				delete(hash_map, val)
			}
			val = val * 2
			if _, ok := hash_map[val]; ok {
				hash_map[val]--
				if hash_map[val] == 0 {
					delete(hash_map, val)
				}
				res = append(res, val/2)
			}
		}
	}
	if len(res) == len(changed)/2 {
		return res
	}
	return nil
}

func main() {
	fmt.Println(find([]int{1, 2, 3, 6, 4, 8}))
}
