/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain
any leading 0's.

Increment the large integer by one and return the resulting array of digits.
*/

package main

import "fmt"

func plus_one(arr []int) []int {
	n := len(arr) - 1
	carry := (arr[n] + 1) / 10
	arr[n] = (arr[n] + 1) % 10
	for i := n - 1; i >= 0; i-- {
		temp := carry
		carry = (arr[i] + carry) / 10
		arr[i] = (arr[i] + temp) % 10
	}

	if carry == 0 {
		return arr
	}
	temp := []int{carry}
	temp = append(temp, arr...)
	return temp
}

func main() {
	// fmt.Println(plus_one([]int{1, 2, 3}))
	fmt.Println(plus_one([]int{8, 9, 9, 9}))
}
