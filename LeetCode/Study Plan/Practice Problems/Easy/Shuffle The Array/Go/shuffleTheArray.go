package main

import "fmt"

func shuffle(nums []int, n int) []int {
	size := 2 * n
	i, j := 0, n

	res := make([]int, size)
	k := 0
	for i < n && j < size {
		res[k] = nums[i]
		res[k+1] = nums[j]
		k += 2
		i++
		j++
	}

	return res
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
