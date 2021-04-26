// Given an integer k and a string s, find the length of the longest substring that contains at most k
// distinct characters.

// For example, given s = "abcba" and k = 2, the longest substring with k distinct characters is "bcb".

package main

import (
	"fmt"
	"math"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	k := 2
	s := "abcba"

	mp := make(map[rune]int)
	x, i, j := 0, 0, 0
	ans := int(math.MinInt32)
	for j, i = 0, 0; j < len(s); j++ {
		mp[rune(s[j])]++
		if mp[rune(s[j])] == 1 {
			fmt.Println("j", string(s[j]), j)
			x++
		}
		for x > k && i <= j {
			fmt.Println("i", string(s[i]), i)
			mp[rune(s[i])]--
			if mp[rune(s[i])] == 0 {
				x--
			}
			i++
		}
		ans = maxInt(ans, j-i+1)
	}
	fmt.Println(ans)
}
