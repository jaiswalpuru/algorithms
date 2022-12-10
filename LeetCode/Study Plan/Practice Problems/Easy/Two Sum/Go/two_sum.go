package main

import "fmt"

func two_sum(arr []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		mp[arr[i]] = i
	}
	res := []int{}
	for i := 0; i < len(arr); i++ {
		val := target - arr[i]
		if _, ok := mp[val]; ok {
			if i != mp[val] {
				res = append(res, []int{mp[val], i}...)
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(two_sum([]int{2, 7, 11, 15}, 9))
}
