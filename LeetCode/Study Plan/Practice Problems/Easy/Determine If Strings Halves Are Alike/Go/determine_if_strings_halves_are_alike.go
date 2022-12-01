package main

import (
	"fmt"
	"strings"
)

func determine_if_strings_halves_are_alike(s string) bool {
	n := len(s)
	cnt := 0
	s = strings.ToUpper(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			cnt++
		}
	}
	for i := len(s) / 2; i < n; i++ {
		if s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			cnt--
		}
	}
	return cnt == 0
}

func main() {
	fmt.Println(determine_if_strings_halves_are_alike("book"))
}
