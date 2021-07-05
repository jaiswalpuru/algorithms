/*
You have n coins and you want to build a staircase with these coins.
The staircase consists of k rows where the ith row has exactly i coins.
The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.

Example 1:
Input: n = 5
Output: 2
Explanation: Because the 3rd row is incomplete, we return 2.

Example 2:
Input: n = 8
Output: 3
Explanation: Because the 4th row is incomplete, we return 3.
*/

package main

import (
	"fmt"
	"math"
)

func compute_nth_sum(n int) int { return n * (n + 1) / 2 }

func num_staircase(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	l, r := 0, n
	for l <= r {
		mid := l + (r-l)/2
		temp := compute_nth_sum(mid)

		if temp == n {
			return mid
		}

		if temp > n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

func num_staircase_math(n int) int { return int(math.Sqrt(float64(2*n)+0.25) - 0.5) }

func main() {
	fmt.Println(num_staircase(8))
	fmt.Println(num_staircase(5))
	fmt.Println(num_staircase_math(8))
	fmt.Println(num_staircase_math(5))
}
