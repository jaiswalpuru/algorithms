package main

import "fmt"

func greatest_common_divisor_of_array(arr []int) int {
	min, max := arr[0], arr[0]
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return gcd(min, max)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Println(greatest_common_divisor_of_array([]int{2, 5, 6, 9, 10}))
}
