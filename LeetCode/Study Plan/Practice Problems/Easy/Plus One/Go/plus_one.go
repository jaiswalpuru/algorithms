package main

import "fmt"

func plus_one(arr []int) []int {
	n := len(arr)
	carry := 1
	for i := n - 1; i >= 0; i-- {
		sum := carry + arr[i]
		arr[i] = sum % 10
		carry = sum / 10
	}
	res := []int{}
	if carry != 0 {
		res = append(res, carry)
	}
	res = append(res, arr...)
	return res
}

func main() {
	fmt.Println(plus_one([]int{1, 2, 3}))
}
