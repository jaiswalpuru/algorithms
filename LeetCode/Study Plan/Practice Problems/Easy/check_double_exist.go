/*
Given an array arr of integers, check if there exists two integers N and M such that N is the double of M
( i.e. N = 2 * M).

More formally check if there exists two indices i and j such that :

i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]
*/

package main

import "fmt"

func exists(arr []int) bool {
	mp := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			mp[arr[i]]++
		} else {
			mp[arr[i]] = i
		}
	}

	for i := 0; i < n; i++ {
		val := arr[i]
		if val == 0 {
			t := mp[val]
			if t > 1 {
				return true
			}
		} else if _, ok := mp[val*2]; ok {
			if mp[val*2] != mp[val] {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(exists([]int{10, 2, 5, 3}))
	fmt.Println(exists([]int{7, 1, 14, 11}))
	fmt.Println(exists([]int{3, 1, 7, 11}))
}
