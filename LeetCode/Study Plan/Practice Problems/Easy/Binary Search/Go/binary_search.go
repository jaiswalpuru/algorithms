package main

import "fmt"

func binary_search(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l+r)/2
		if arr[mid] == target {
			return mid+1
		}else if arr[mid] > target {
			r = mid-1
		}else{
			l= mid+1
		}
	}
	return -1
}

func binary_search_recursive(arr []int, target int) int {
	return _bsearch(arr, target, 0, len(arr)-1)
}

func _bsearch(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right)/2
	if arr[mid] == target {
		return mid
	}else if arr[mid] > target {
		return _bsearch(arr, target, left, mid-1)
	}else {
		return _bsearch(arr, target, mid+1, right)
	}
}


func main() {
	fmt.Println(binary_search_recursive([]int{-1,0,3,5,9,12}, 9))
}