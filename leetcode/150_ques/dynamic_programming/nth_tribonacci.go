/*
The Tribonacci sequence Tn is defined as follows:

T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

Given n, return the value of Tn.

Example 1:
Input: n = 4
Output: 4
Explanation:
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4

Example 2:
Input: n = 25
Output: 1389537
*/

package main

import "fmt"

func nth_tribonacci_dp(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}
	arr := make([]int, n+1)
	arr[0], arr[1], arr[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
	}
	return arr[n]
}

func nth_tribonacci_recursive(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return nth_tribonacci_recursive(n-1) + nth_tribonacci_recursive(n-2) + nth_tribonacci_recursive(n-3)
}

func main() {
	fmt.Println(nth_tribonacci_dp(4))
	fmt.Println(nth_tribonacci_recursive(4))
	fmt.Println(nth_tribonacci_dp(25))
	fmt.Println(nth_tribonacci_recursive(25))
}
