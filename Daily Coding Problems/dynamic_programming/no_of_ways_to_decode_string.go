/*
Given mapping a=1, b=2, ...., z=26 and an ecoded message count the number of ways it can be decoded

eg : I/P : "111"  O/P : 3 -> "aaa", "ka", "ak"
*/

package main

import (
	"fmt"
	"strconv"
)

func num_ways_to_decode(str string, total int) int {

	//for empty and single letter string
	if len(str) <= 1 {
		return 1
	}

	//no message starts with zero
	if str[0] == '0' {
		return 0
	}

	total += num_ways_to_decode(str[1:], 0)
	val, _ := strconv.Atoi(str[:2])
	if val <= 26 {
		total += num_ways_to_decode(str[2:], 0)
	}

	return total

}

func num_ways_to_decode_dp(str string) int {
	dp := make([]int, len(str)+1)
	dp[len(str)] = 1
	n := len(str) - 1
	for i := n; i >= 0; i-- {
		if str[i] == '0' {
			dp[i] = 0
		} else if i == len(str)-1 {
			dp[i] = 1
		} else {
			dp[i] += dp[i+1]
			fmt.Println(i, i+2)
			val, _ := strconv.Atoi(str[i : i+2])
			if val <= 26 {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}

func main() {
	fmt.Println(num_ways_to_decode("111", 0))
	fmt.Println(num_ways_to_decode_dp("111"))
	fmt.Println(num_ways_to_decode_dp("001"))
}
