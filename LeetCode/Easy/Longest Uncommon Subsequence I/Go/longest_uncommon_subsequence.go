package main

import "fmt"

func longest_uncommon_subsequence(a, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	} else {
		return len(b)
	}
}

func main() {
	fmt.Println(longest_uncommon_subsequence("aba", "cdc"))
}
