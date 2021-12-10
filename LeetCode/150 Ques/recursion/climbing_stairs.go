/*

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps


Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

*/

package main

import "fmt"

func climbing_stairs(n int) int {

	if n == 1 || n == 0 || n == 2 {
		return n
	}

	return climbing_stairs(n-1) + climbing_stairs(n-2)
}

//non recursive
func climbing_stairs_non_recursive(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	arr := make([]int, n+1)
	arr[0], arr[1], arr[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func main() {
	fmt.Println(climbing_stairs(2))
	fmt.Println(climbing_stairs(3))
}