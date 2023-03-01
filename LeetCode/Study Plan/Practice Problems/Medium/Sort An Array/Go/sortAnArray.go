package main

import "fmt"

//---------------------------- applying merge sort ----------------------------
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := len(nums) / 2
	left := sortArray(nums[:pivot])
	right := sortArray(nums[pivot:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	n := len(left)
	m := len(right)
	arr := make([]int, n+m)
	l, r, k := 0, 0, 0

	for l < n && r < m {
		if left[l] < right[r] {
			arr[k] = left[l]
			l++
		} else {
			arr[k] = right[r]
			r++
		}
		k++
	}

	for l < n {
		arr[k] = left[l]
		l++
		k++
	}

	for r < m {
		arr[k] = right[r]
		r++
		k++
	}

	return arr
}

//-----------------------------------------------------------------------------

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
}
