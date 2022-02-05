package main

import "fmt"

func remove(arr []int) int {
	n := len(arr)
	l, r := 0, n-1

	for l < r && arr[l] <= arr[l+1] {
		l++
	}

	//if l reaches the end of the array then it means that whole array is sorted
	if l == n-1 {
		return 0
	}

	for r > 0 && arr[r] >= arr[r-1] {
		r--
	}

	remove := min(n-l-1, r)

	for i := 0; i < l+1; i++ {
		if arr[i] <= arr[r] {
			remove = min(remove, r-i-1)
		} else if r < n-1 {
			r++
		} else {
			break
		}
	}

	return remove
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(remove([]int{1, 2, 3, 10, 4, 2, 3, 5}))
}
