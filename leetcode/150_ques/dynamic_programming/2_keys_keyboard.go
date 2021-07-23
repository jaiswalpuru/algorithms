/*
There is only one character 'A' on the screen of a notepad. You can perform two operations on this
notepad for each step:

Copy All: You can copy all the characters present on the screen (a partial copy is not allowed).
Paste: You can paste the characters which are copied last time.
Given an integer n, return the minimum number of operations to get the character 'A' exactly n times on the screen.

Example 1:
Input: n = 3
Output: 3
Explanation: Intitally, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'.

Example 2:
Input: n = 1
Output: 0
*/

package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func two_key_keyboards(n int) int {
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 0
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			res[i] = 2 + res[i/2]
		} else {
			res[i] = i
			for j := int(math.Sqrt(float64(i))); j > 1; j-- {
				//fmt.Println(i, j)
				if i%j == 0 {
					res[i] = min(res[i], res[i/j]+j)
				}
			}
		}
	}
	return res[n]
}

func main() {
	fmt.Println(two_key_keyboards(3))
	fmt.Println(two_key_keyboards(1))
	fmt.Println(two_key_keyboards(9))
}
