/*
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array.
*/

package main

import "fmt"

func majority_element(arr []int) int {
	count_num := make(map[int]int)
	n := len(arr)
	ans, cnt := 0, 0
	for i := 0; i < n; i++ {
		count_num[arr[i]]++
		if count_num[arr[i]] > cnt {
			cnt = count_num[arr[i]]
			ans = arr[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(majority_element([]int{6, 5, 5}))
}
