package main

import "fmt"

//------------------------------using binary search----------------------------
func longest_repeating_character_replacement_binary_search(s string, k int) int {
	l, r := 1, len(s)+1
	for l+1 < r {
		mid := l + (r-l)/2
		if valid(s, mid, k) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

func valid(s string, substring_len, k int) bool {
	freq := make(map[byte]int)
	start := 0
	max_freq := 0
	for end := 0; end < len(s); end++ {
		freq[s[end]]++
		if end-start+1 > substring_len {
			freq[s[start]]--
			start++
		}
		max_freq = max(max_freq, freq[s[end]])
		if substring_len-max_freq <= k {
			return true
		}
	}
	return false
}

//------------------------------using binary search----------------------------

func longest_repeating_character_replacement(s string, k int) int {
	win_start := 0
	max_freq := 0
	max_length := 0
	freq := make(map[byte]int)
	for win_end := 0; win_end < len(s); win_end++ {
		freq[s[win_end]]++
		max_freq = max(max_freq, freq[s[win_end]])
		for win_end-win_start-max_freq+1 > k {
			freq[s[win_start]]--
			win_start++
		}
		max_length = max(max_length, win_end-win_start+1)
	}
	return max_length
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
