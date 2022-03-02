package main

import "fmt"

func is_subsequence(s, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if t[j] == s[i] {
			i++
			j++
		} else {
			j++
		}
	}
	fmt.Println(i, j)
	if i == n {
		return true
	}
	return false
}

func main() {
	fmt.Println(is_subsequence("abc", "ahbcdg"))
}
