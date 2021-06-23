/*
	Given an array of numbers of length n, find both the min and max using less than 2*(n-2) comparisons
*/

package main

import "fmt"

//compares and return the max and min value of the pair
func compare(a, b int) (int, int) {
	if b > a {
		return a, b
	}
	return b, a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min_max(arr []int) (int, int) {
	min_ele, max_ele := arr[0], arr[0]

	//make the list odd so that pair can be formed easily
	if len(arr)%2 == 0 {
		arr = append(arr, arr[len(arr)-1])
	}

	for i := 1; i < len(arr); i += 2 {
		smaller, bigger := compare(arr[i], arr[i+1])
		min_ele = min(min_ele, smaller)
		max_ele = max(max_ele, bigger)
	}
	return min_ele, max_ele
}

//using divide and conquer
func min_max_using_divide_and_conquer(arr []int) (int, int) {
	if len(arr) == 0 {
		panic("No elements in the array")
	}
	if len(arr) == 1 {
		return arr[0], arr[0]
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[1], arr[0]
		} else {
			return arr[0], arr[1]
		}
	}

	n := len(arr) / 2
	lmin, lmax := min_max_using_divide_and_conquer(arr[:n])
	rmin, rmax := min_max_using_divide_and_conquer(arr[n:])
	return min(lmin, rmin), max(lmax, rmax)
}

func main() {
	arr := []int{4, 2, 7, 5, -1, 3, 6}
	fmt.Println(min_max(arr))
	fmt.Println(min_max_using_divide_and_conquer(arr))
}
