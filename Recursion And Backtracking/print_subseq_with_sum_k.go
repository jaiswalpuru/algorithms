//print subsequence whose sum is k

package main

import "fmt"

func print_sub_k(arr []int, k int) {
	sum := 0
	_print(arr, k, []int{}, 0, &sum)
}

func _print(arr []int, target int, temp []int, start int, sum *int) {
	if start >= len(arr) {
		if target == *sum {
			fmt.Println(temp)
		}
		return
	}

	t := append(temp, arr[start])
	*sum += arr[start]
	_print(arr, target, t, start+1, sum)
	*sum -= arr[start]
	t = t[:len(t)-1]
	_print(arr, target, t, start+1, sum)
}

func main() {
	print_sub_k([]int{3, 2, 1, 6, 5, 4}, 6)
}
