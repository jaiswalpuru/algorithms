package main

import "fmt"

func longest_repeating_character_replacement(s string, k int) int {
	l, max_freq := 0, 0
	mp := make(map[byte]int)
	res := 0
	for r := 0; r < len(s); r++ {
		mp[s[r]]++
		max_freq = max(max_freq, mp[s[r]])
		if r-l+1-max_freq > k {
			mp[s[l]]--
			l++
		}
		res = r - l + 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longest_repeating_character_replacement("ABAB", 2))
}
