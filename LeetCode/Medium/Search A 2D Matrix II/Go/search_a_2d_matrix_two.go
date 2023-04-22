package main

import "fmt"

//-------------------generic method---------------------
func search(arr []int, l, r, target int) bool {
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return true
		}
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func search_two_d_matrix_two(arr [][]int, target int) bool {
	n, m := len(arr), len(arr[0])
	for i := 0; i < n; i++ {
		if search(arr[i], 0, m-1, target) {
			return true
		}
	}
	return false
}

//-------------------generic method---------------------

//-------------------divide and conquer method---------------------
func search_two_d_matrix_two_dc(arr [][]int, target int) bool {
	return search(0, len(arr[0])-1, 0, len(arr)-1, arr, target)
}

func search(left, right, up, down int, arr [][]int, target int) bool {
	if left > right || up > down {
		return false
	} else if target < arr[up][left] || target > arr[down][right] {
		return false
	}

	row := up
	mid := (left + right) / 2
	for row <= down && arr[row][mid] <= target {
		if arr[row][mid] == target {
			return true
		}
		row++
	}

	return search(left, mid-1, row, down, arr, target) || search(mid+1, right, up, row-1, arr, target)
}

//-------------------divide and conquer method---------------------

func main() {
	fmt.Println(search_two_d_matrix_two([][]int{
		{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30},
	}, 5))
}
