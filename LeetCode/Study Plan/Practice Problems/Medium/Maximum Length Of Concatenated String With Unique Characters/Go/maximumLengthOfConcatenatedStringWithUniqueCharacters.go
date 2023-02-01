package main

import (
	"fmt"
	"strings"
)

func maxLength(arr []string) int {
	maxLength := 0
	backtrack(arr, &maxLength, 0, []string{})
	return maxLength
}

func backtrack(arr []string, maxLength *int, index int, set []string) {
	l := 0
	for i := index; i < len(arr); i++ {
		set = append(set, arr[i])
		if isUnique(strings.Join(set, "")) {
			l = len(strings.Join(set, ""))
			backtrack(arr, maxLength, i+1, set)
		}
		*maxLength = max(*maxLength, l)
		set = set[:len(set)-1]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isUnique(s string) bool {
	size := len(s)
	hashMap := make(map[byte]int)
	for i := 0; i < size; i++ {
		hashMap[s[i]]++
		if hashMap[s[i]] > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
}
