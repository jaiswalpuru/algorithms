package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longest_substring_without_repeating_characters(s string) int {
	cnt_map := make(map[byte]int)
	str := []byte(s)
	i, j, n := 0, 0, len(str)
	l := 0
	for j < n {
		if _, ok := cnt_map[str[j]]; ok {
			l = max(l, j-i)
			i = max(i, cnt_map[str[j]]+1)
			cnt_map[str[j]] = j
		} else {
			cnt_map[str[j]] = j
		}
		j++
	}
	l = max(l, j-i)
	return l
}

func main() {
	fmt.Println(longest_substring_without_repeating_characters("abcabcbb"))
}
