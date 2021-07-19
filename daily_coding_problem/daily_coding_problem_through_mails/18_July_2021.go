/*
Given a number represented by a list of digits, find the next greater permutation of a number,
in terms of lexicographic ordering. If there is not greater permutation possible, return the
permutation with the lowest value/ordering.

For example, the list [1,2,3] should return [1,3,2]. The list [1,3,2] should return [2,1,3].
The list [3,2,1] should return [1,2,3].
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

func next_permutation(arr []int) []int {
	i := len(arr) - 2
	f := true
	for i >= 0 && arr[i+1] <= arr[i] {
		i--
	}
	if i >= 0 {
		f = false
		j := len(arr) - 1
		for arr[j] <= arr[i] {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
		reverse(&arr, i+1)
	}
	if f {
		reverse(&arr, 0)
	}
	return arr
}

func main() {
	fmt.Println(next_permutation([]int{1, 2, 3}))
	fmt.Println(next_permutation([]int{3, 2, 1}))
}
