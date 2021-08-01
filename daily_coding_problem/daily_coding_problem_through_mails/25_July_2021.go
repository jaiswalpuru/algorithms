/*
Given a list of integers and a number K, return which contiguous elements of the list sum to K.

For example, if the list is [1, 2, 3, 4, 5] and K is 9, then it should return [2, 3, 4], since 2 + 3 + 4 = 9.
*/

package main

import "fmt"

func contiguous_elements(arr []int, target int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}

	l, r, sum := 0, 0, 0
	ele := make([]int, 0)
	for r < n {
		if sum+arr[r] == target || r >= n {
			ele = append(ele, arr[r])
			break
		}

		if sum+arr[r] > target {
			sum -= arr[l]
			ele = ele[l+1:]
			l++
		} else {
			sum += arr[r]
			ele = append(ele, arr[r])
			r++
		}
	}
	return ele
}

func main() {
	fmt.Println(contiguous_elements([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(contiguous_elements([]int{3, 3, 3, 1, 2}, 9))
	fmt.Println(contiguous_elements([]int{-1, 0, 1, 0, 0}, 0))
}
