package main

import "fmt"

func construct_k_palindrome_strings(s string, k int) bool {
	if len(s) < k {
		return false
	}
	mp := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		mp[s[i]]++
	}
	count_odd := 0
	for _, v := range mp {
		if v%2 == 1 {
			if count_odd != 0 {
				return false
			}
			count_odd = v
		}
	}

	if count_odd > k {
		return false
	}

	return true
}

func main() {
	fmt.Println(construct_k_palindrome_strings("annabelle", 2))
}
