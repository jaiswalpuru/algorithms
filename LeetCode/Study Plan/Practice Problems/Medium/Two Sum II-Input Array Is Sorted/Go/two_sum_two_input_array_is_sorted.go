package main

import "fmt"

func two_sum(arr []int, target int) []int {
	i, j := 0, len(arr)-1
	for i < j {
		sum := arr[i] + arr[j]
		if sum == target {
			return []int{i+1, j+1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

func main() {
	fmt.Println(two_sum([]int{2, 7, 11, 15}, 9))
}
