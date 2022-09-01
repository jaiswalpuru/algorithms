package main

import "fmt"

//-----------------------brute force---------------------
func sum_of_all_odd_length_subarrays(arr []int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (j-i+1)%2 != 0 {
				sum := 0
				for k := i; k <= j; k++ {
					sum += arr[k]
				}
				res += sum
			}
		}
	}
	return res
}

//-----------------------brute force---------------------

//-----------------------brute force---------------------
func sum_of_all_odd_length_subarrays_better(arr []int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		curr := 0
		for j := i; j < n; j++ {
			curr += arr[j]
			if (j-i)%2 == 0 {
				res += curr
			}
		}
	}
	return res
}

//-----------------------brute force---------------------

func main() {
	fmt.Println(sum_of_all_odd_length_subarrays_better([]int{1, 4, 2, 5, 3}))
}
