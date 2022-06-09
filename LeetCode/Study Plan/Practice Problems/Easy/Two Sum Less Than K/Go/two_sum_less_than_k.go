package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func two_sum_less_than_k(arr []int, target int) int {
	sum := -1
	for i := 0; i < len(arr); i++ {
		temp := -1
		for j := i + 1; j < len(arr); j++ {
			if target > arr[i]+arr[j] {
				temp = max(temp, arr[j]+arr[i])
			}
		}
		sum = max(sum, temp)
	}
	return sum
}

func main() {
	fmt.Println(two_sum_less_than_k([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60))
}
