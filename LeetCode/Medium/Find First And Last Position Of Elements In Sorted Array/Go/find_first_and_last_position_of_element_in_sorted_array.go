package main

import "fmt"

func find_first_and_last_position_of_element_in_sorted_array(arr []int, target int) []int {
	res := make([]int, 2)
	res[0] = find_lower_bound(arr, target)
	res[1] = res[0]
	res[0] = find_upper_bound(arr, target)
	return res
}

func find_lower_bound(arr []int, target int) int {
	pos := -1
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			pos = mid
			l = mid + 1
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return pos
}

func find_upper_bound(arr []int, target int) int {
	pos := -1
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			pos = mid
			r = mid - 1
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return pos
}

func main() {
	fmt.Println(find_first_and_last_position_of_element_in_sorted_array([]int{5, 7, 7, 8, 8, 10}, 8))
}
