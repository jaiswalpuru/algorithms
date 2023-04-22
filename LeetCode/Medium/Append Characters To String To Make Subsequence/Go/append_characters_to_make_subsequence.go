package main

import "fmt"

func append_characters_to_make_subsequence(s, t string) int {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return len(t) - j
}

func main() {
	fmt.Println(append_characters_to_make_subsequence("coaching", "coding"))
}
