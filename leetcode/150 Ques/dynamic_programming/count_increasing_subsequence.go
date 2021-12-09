/*
This is not a part of this series. It was asked in google onsite interview.
Given an array arr of size n. Find the number of triples (i, j, k) where:

i < j < k
arr[i] < arr[j] < arr[k]

Example 1:
Input: arr = [1, 2, 3, 4, 5]
Output: 10
Explanation:
1. 1 2 3
2. 1 2 4
3. 1 2 5
4. 1 3 4
5. 1 3 5
6. 1 4 5
7. 2 3 4
8. 2 3 5
9. 2 4 5
10. 3 4 5

Example 2:
Input: arr = [1, 2, 3, 5, 4]
Output: 7

Example 3:
Input: arr = [5, 4, 3, 2, 1]
Output: 0
*/

package main

import "fmt"

func count_increasing_subsequence(arr []int, k int) int {
	dp := make([][]int, k)
	n := len(arr)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for t := 1; t < k; t++ {
		for i := t; i < n; i++ {
			for j := t - 1; j < i; j++ {
				if arr[j] < arr[i] {
					dp[t][i] += dp[t-1][j]
				}
			}
		}
	}

	ans := 0
	for i := k - 1; i < n; i++ {
		ans += dp[k-1][i]
	}

	return ans
}

func main() {
	fmt.Println(count_increasing_subsequence([]int{1, 2, 3, 4, 5}, 3))
}
