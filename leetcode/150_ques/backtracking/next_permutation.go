/*
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such an arrangement is not possible, it must rearrange it as the lowest possible order
(i.e., sorted in ascending order).

The replacement must be in place and use only constant extra memory.

Example 1:
Input: nums = [1,2,3]
Output: [1,3,2]

Example 2:
Input: nums = [3,2,1]
Output: [1,2,3]

Example 3:
Input: nums = [1,1,5]
Output: [1,5,1]

Example 4:
Input: nums = [1]
Output: [1]
*/

package main

import "fmt"

func reverse(a *[]int, i int) {
	m := len((*a)[i:])
	n := len((*a)[i:]) / 2
	for j := i; j < n; j++ {
		(*a)[j], (*a)[m-1-j] = (*a)[m-1-j], (*a)[j]
	}
}

func next_permute(arr []int) []int {
	i := len(arr) - 2

	//Find the first decreasing element
	for i >= 0 && arr[i+1] <= arr[i] {
		i--
	}
	if i >= 0 {
		j := len(arr) - 1
		for arr[j] <= arr[i] {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
		reverse(&arr, i+1)
	}
	return arr
}

func main() {
	fmt.Println(next_permute([]int{1, 2, 3}))
	fmt.Println(next_permute([]int{3, 2, 1}))
	fmt.Println(next_permute([]int{1, 1, 5}))
	fmt.Println(next_permute([]int{1}))
	fmt.Println(next_permute([]int{1, 5, 8, 4, 7, 6, 5, 3, 1}))
}
