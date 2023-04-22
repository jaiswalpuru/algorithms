/*
You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately,
due to some error, one of the numbers in s got duplicated to another number in the set, which results
in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.
*/

package main

import "fmt"

func find_err(arr []int) []int {
	ans := make([]int, 2)
	n := len(arr)

	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[arr[i]]++
		if mp[arr[i]] == 2 {
			ans[0] = arr[i]
		}
	}

	for i := 1; i <= n; i++ {
		if _, ok := mp[i]; !ok {
			ans[1] = i
		}
	}

	return ans

}

func main() {
	fmt.Println(find_err([]int{1, 2, 2, 4}))
	fmt.Println(find_err([]int{1, 1}))
}
