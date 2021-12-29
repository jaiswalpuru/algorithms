/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/

package main

import "fmt"

func _pow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, n)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		half_ans := pow(x, n/2)
		return half_ans * half_ans
	} else {
		half_ans := pow(x, n/2)
		return half_ans * half_ans * x
	}
}

func main() {
	fmt.Println(_pow(2.00000, 10))
}
