// There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time.
// Given N, write a function that returns the number of unique ways you can climb the staircase.
// The order of the steps matters.
// For example, if N is 4, then there are 5 unique ways:
// 1, 1, 1, 1
// 2, 1, 1
// 1, 2, 1
// 1, 1, 2
// 2, 2
// What if, instead of being able to climb 1 or 2 steps at a time,
// you could climb any number from a set of positive integers X? For example,
// if X = {1, 3, 5}, you could climb 1, 3, or 5 steps at a time.

package main

import "fmt"

func main() {
	var N = 10 //eg
	//X := []int{1, 2} //one can climb either 1 or 2 steps at a time
	dp := make([]int, N+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp[N])

	//Now if X = {1, 3, 5} we can climb 1 3 or 5 steps at a time
	dp = make([]int, N+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2
	dp[4] = 3
	dp[5] = 5
	for i := 6; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp[N])
}
