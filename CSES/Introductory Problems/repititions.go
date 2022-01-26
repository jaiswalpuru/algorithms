package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func repititions(s string) int {
	n := len(s)
	res := -1
	length := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			length++
		} else {
			res = max(res, length)
			length = 1
		}
	}
	return res
}

func main() {
	fmt.Println(repititions("ATTCGGGA"))
}
