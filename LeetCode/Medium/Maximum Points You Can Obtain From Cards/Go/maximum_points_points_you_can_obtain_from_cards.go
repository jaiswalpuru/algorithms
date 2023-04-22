package main

import "fmt"

//--------------------------brute force--------------------------------
func max_points_you_can_obtain_from_cards(arr []int, k int) int {
	n := len(arr)
	total := 0
	_recur(arr, &total, k, 0, n-1, 0)
	return total
}

func _recur(arr []int, total *int, k, i, n int, temp_sum int) {
	if k == 0 {
		*total = max(*total, temp_sum)
		return
	}

	t := temp_sum + arr[n]
	_recur(arr, total, k-1, i, n-1, t)
	t += arr[i] - arr[n]
	_recur(arr, total, k-1, i+1, n, t)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//--------------------------brute force--------------------------------

//-------------------------------eff approach--------------------------
func max_points_you_can_obtain_from_cards_eff(arr []int, k int) int {
	front := make([]int, k+1)
	back := make([]int, k+1)
	n := len(arr)

	for i := 0; i < k; i++ {
		front[i+1] = arr[i] + front[i]
		back[i+1] = arr[n-i-1] + back[i]
	}

	total := 0
	for i := 0; i <= k; i++ {
		sum := front[i] + back[k-i]
		total = max(total, sum)
	}
	return total
}

//-------------------------------eff approach--------------------------

func main() {
	fmt.Println(max_points_you_can_obtain_from_cards_eff([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}
