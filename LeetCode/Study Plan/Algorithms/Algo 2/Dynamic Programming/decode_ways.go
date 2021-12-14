/*
A message containing letters from A-Z can be encoded into numbers using the following mapping:

'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping
above (there may be multiple ways). For example, "11106" can be mapped into:

"AAJF" with the grouping (1 1 10 6)
"KJF" with the grouping (11 10 6)
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".

Given a string s containing only digits, return the number of ways to decode it.

The answer is guaranteed to fit in a 32-bit integer.
*/

package main

import (
	"fmt"
	"strconv"
)

func decode_ways(s string) int {
	n := len(s) - 1
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	for i := n; i >= 0; i-- {
		if i == n {
			dp[i] = 1
		} else if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] += dp[i+1]
			val, _ := strconv.Atoi(s[i : i+2])
			if val <= 26 {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}

func main() {
	fmt.Println(decode_ways("226"))
}
