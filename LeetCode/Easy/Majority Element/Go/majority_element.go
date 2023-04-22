package main

import "fmt"

//O(1) space soln would be to sort the elements and return arr[n/2]

func majority_element(arr []int) int {
	mp := make(map[int]int)
	n := len(arr)
	val := 0
	for i := 0; i < n; i++ {
		mp[arr[i]]++
		if mp[val] < mp[arr[i]] {
			val = arr[i]
		}
	}
	return val
}

func main() {
	fmt.Println(majority_element([]int{3, 2, 3}))
}
