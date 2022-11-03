package main

import "fmt"

func longest_palindrome(s string) int {
	n := len(s)
	cnt := make(map[byte]int)
	for i := 0; i < n; i++ {
		cnt[s[i]]++
	}
	odd := 0
	l := 0
	for _, v := range cnt {
		if v%2 == 0 {
			l += v
		} else {
			if odd < v {
				if odd == 0 {
					odd = v
				} else {
					l += (odd - 1)
					odd = v
				}
			} else {
				l += (v - 1)
			}
		}
	}
	l += odd
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longest_palindrome("abccccdd"))
}
