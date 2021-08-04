/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
*/

package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(arr []int) int {
	mp := make(map[int]int)
	n := len(arr)

	for v := range arr {
		mp[arr[v]] = 1
	}

	res := 0
	for i := 0; i < n; i++ {
		if _, ok := mp[arr[i]-1]; !ok {
			curr := arr[i]

			temp := 0
			_, ok := mp[curr]
			for ok {
				temp += 1
				curr = curr + 1
				_, ok = mp[curr]
			}

			res = max(res, temp)
		}
	}

	return res
}

func main() {
	fmt.Println(lcs([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(lcs([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
