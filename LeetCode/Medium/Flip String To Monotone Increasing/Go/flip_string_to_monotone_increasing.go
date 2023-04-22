package main

import "fmt"

//For my future reference
/* - refer LC
But can we go above and beyond? Yes, but we need to think of the problem from a different perspective. Let us iterate through the array from left to right. Suppose we know that s[0...i - 1] is a monotone increasing array, if s[i] is 1 then we maintain the monotone increasing property and don't have to do anything. However, what if s[i] is 0? Well, we have two choice, we can either do flip that element or flips every 1 in s[0...i]. Both of these action maintain the monotone increasing property. Therefore, we get the following recurrence relationship, let T(i) be the minimum number of flips needed to make s[0...i] a monotone increasing array
*/

//-------------------dp---------------------
func minimum_flips_to_monotone_string_dp(s string) int {
	dp := make([]int, len(s)+1)
	ones := 0
	for i := 1; i <= len(s); i++ {
		if s[i-1] == '1' {
			ones++
			dp[i] = dp[i-1]
		} else {
			dp[i] = min(dp[i-1]+1, ones)
		}
	}
	return dp[len(s)]
}

//-------------------dp---------------------

//----------------more space optimized dp------------------
func minimum_flips_to_monotone_string_dp_opt(s string) int {
	ones := 0
	dp := 0
	for i := 1; i <= len(s); i++ {
		if s[i-1] == '1' {
			ones++
		} else {
			dp = min(dp+1, ones)
		}
	}
	return dp
}

//----------------more space optimized dp------------------

//---------------dynamic window concept refer LC------------------
func minimum_flips_to_monotone_string(s string) int {
	zero := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zero++
		}
	}
	res := zero
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zero--
			res = min(res, zero)
		} else {
			zero++
		}
	}
	return res
}

//---------------dynamic window concept refer LC------------------

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimum_flips_to_monotone_string_dp("00110"))
}
